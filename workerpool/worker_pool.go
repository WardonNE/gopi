package workerpool

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/wardonne/gopi/support/maps"
	"github.com/wardonne/gopi/workerpool/driver"
	"github.com/wardonne/gopi/workerpool/job"
)

type WorkerPoolStatus int

// WorkerPoolStatus enums
const (
	WorkerPoolStatusRunning WorkerPoolStatus = iota + 1
	WorkerPoolStatusStopped
)

// WorkerPool is a struct to manage workers
// it accepts Job and push the job to workers
type WorkerPool struct {
	id        uuid.UUID
	name      string
	status    WorkerPoolStatus
	createdAt time.Time
	startAt   time.Time
	stoppedAt time.Time

	mu                 sync.Mutex
	workers            *maps.SyncHashMap[uuid.UUID, *Worker]
	stopChannel        chan struct{}
	watcherStopChannel chan struct{}

	driver     driver.DriverInterface
	maxWorkers int
	// worker configs
	workerConfigs struct {
		batch          int
		maxIdleTime    time.Duration
		maxStoppedTime time.Duration
	}
	// job configs
	jobConfigs struct {
		maxAttempts    int
		retryDelay     time.Duration
		retryMaxDelay  time.Duration
		retryDelayStep time.Duration
		maxExecuteTime time.Duration
	}
}

// DefaultWorkerPool creates WorkerPool instance with specific driver and default configs
// example:
//
//	wp := DefaultWorkerPool(driver.NewQueueDriver())
//	wp.Start()
func DefaultWorkerPool(driver driver.DriverInterface) *WorkerPool {
	wp := new(WorkerPool)
	// basic attributes
	wp.id = uuid.New()
	wp.status = WorkerPoolStatusStopped
	wp.createdAt = time.Now()
	wp.stoppedAt = time.Now()
	// worker container
	wp.workers = maps.NewSyncHashMap[uuid.UUID, *Worker]()

	wp.driver = driver

	// stop signal channel
	wp.stopChannel = make(chan struct{})
	// watcher stop signal channel
	wp.watcherStopChannel = make(chan struct{})
	// configs
	wp.workerConfigs.batch = DefaultWorkerBatch
	wp.workerConfigs.maxIdleTime = DefaultWorkerMaxIdleTime
	wp.workerConfigs.maxStoppedTime = DefaultWorkerStoppedTimeout
	wp.jobConfigs.maxAttempts = DefaultJobMaxAttempts
	wp.jobConfigs.retryDelay = DefaultJobRetryDelay
	wp.jobConfigs.retryMaxDelay = DefaultJobRetryMaxDelay
	wp.jobConfigs.retryDelayStep = DefaultJobRetryDelayStep
	wp.jobConfigs.maxExecuteTime = DefaultJobMaxExecuteTimeTotal
	return wp
}

// NewWorkerPool creates a [WorkerPool] instance with specific driver and custom [Option]s
//
// example:
//
//	wp := NewWorkerPool(driver.NewQueueDriver(), MaxWorkers(10))
//	wp.Start()
func NewWorkerPool(driver driver.DriverInterface, options ...Option) *WorkerPool {
	wp := DefaultWorkerPool(driver)
	for _, option := range options {
		option(wp)
	}
	return wp
}

// NewWorkerWithConfigs creates a [WorkerPool] instance with specific driver and custom [WorkerPoolConfigs].
//
// In fact, it change configs to a list of [Option]s and calls [NewWorkerPool] to create WorkerPool
//
// example:
//
//	wp := NewWorkerWithConfigs(driver.NewQueueDriver(), &WorkerPoolConfigs{
//		MaxWorkers: 10
//	})
//	wp.Start()
func NewWorkerWithConfigs(driver driver.DriverInterface, configs *WorkerPoolConfigs) *WorkerPool {
	return NewWorkerPool(driver, configs.ToOptions()...)
}

// ID returns the unique id of the WorkerPool
func (wp *WorkerPool) ID() uuid.UUID {
	return wp.id
}

// Name returns the name of the WorkerPool if it's added into WorkerPoolManager
//
// It will return empty string if this WorkerPool instance is not added into WorkerPoolManager
func (wp *WorkerPool) Name() string {
	return wp.name
}

// Status returns the active status of the WorkerPool
func (wp *WorkerPool) Status() WorkerPoolStatus {
	return wp.status
}

// CreatedAt returns the created time of the WorkerPool
func (wp *WorkerPool) CreatedAt() time.Time {
	return wp.createdAt
}

// StartedAt returns the started time of the WorkerPool
func (wp *WorkerPool) StartedAt() time.Time {
	return wp.startAt
}

// StoppedAt returns the stopped time of the WorkerPool
func (wp *WorkerPool) StoppedAt() time.Time {
	return wp.stoppedAt
}

// IsRunning returns whether the WorkerPool is running
func (wp *WorkerPool) IsRunning() bool {
	return wp.status == WorkerPoolStatusRunning
}

// IsStopped returns whether the WorkerPool is stopped
func (wp *WorkerPool) IsStopped() bool {
	return wp.status == WorkerPoolStatusStopped
}

func (wp *WorkerPool) stop() {
	wp.status = WorkerPoolStatusStopped
	wp.stoppedAt = time.Now()
}

func (wp *WorkerPool) start() {
	wp.status = WorkerPoolStatusRunning
	wp.startAt = time.Now()
}

// Dispatch dispatches job
func (wp *WorkerPool) Dispatch(job job.JobInterface) bool {
	if wp.IsStopped() {
		return false
	}
	ok := wp.driver.Enqueue(job)
	wp.spawnWorkers()
	return ok
}

// Start starts the workerpool
func (wp *WorkerPool) Start() {
	wp.start()
	wp.spawnWorkers()
	go func() {
		for {
			select {
			case <-wp.watcherStopChannel:
				return
			default:
				workers := wp.workers.Values()
				for _, w := range workers {
					if w.ShouldStop() {
						w.Stop()
					} else if w.ShouldRelease() {
						wp.workers.Remove(w.id)
						w.Release()
					}
				}
				time.Sleep(time.Second)
			}
		}
	}()
}

// Stop stops the worker pool and all the workers
func (wp *WorkerPool) Stop() {
	wp.workers.Range(func(entry *maps.Entry[uuid.UUID, *Worker]) bool {
		entry.Value.Stop()
		return true
	})
	wp.stop()
}

// Release releases and removes the workerpool from the [WorkerPoolManager]
func (wp *WorkerPool) Release() {
	// if the worker pool is running, stop it first
	if wp.IsRunning() {
		wp.stop()
	}
	// notify watcher to stop
	wp.watcherStopChannel <- struct{}{}
	// release workers
	wp.workers.Range(func(entry *maps.Entry[uuid.UUID, *Worker]) bool {
		entry.Value.Release()
		return true
	})
	wp.workers.Clear()
	close(wp.stopChannel)
	close(wp.watcherStopChannel)
}

func (wp *WorkerPool) spawnWorkers() {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	if wp.workers.Count() >= wp.maxWorkers {
		return
	}
	if wp.driver.IsEmpty() {
		return
	}
	nc := wp.maxWorkers / wp.workerConfigs.batch
	if nc == 0 {
		nc = wp.maxWorkers
	}
	if count := wp.driver.Count(); nc > count {
		nc = count
	}
	c := 0
	// awake sleeping workers
	wp.workers.Range(func(entry *maps.Entry[uuid.UUID, *Worker]) bool {
		if entry.Value.IsStopped() {
			go entry.Value.Start()
			c++
		}
		return true
	})
	// create some new workers
	for i := c; i < nc; i++ {
		if wp.workers.Count() >= wp.maxWorkers {
			return
		}
		w := hire(wp)
		wp.workers.Set(w.id, w)
		go w.Start()
	}
}

// Workers returns a slice of Workers
func (wp *WorkerPool) Workers() []*Worker {
	return wp.workers.Values()
}
