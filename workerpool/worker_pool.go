package workerpool

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/wardonne/gopi/support/maps"
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
	status    WorkerPoolStatus
	createdAt time.Time
	startAt   time.Time
	stoppedAt time.Time

	mu          sync.Mutex
	workers     *maps.HashMap[uuid.UUID, *Worker]
	stopChannel chan struct{}

	driver        IWorkerPoolDriver
	maxWorkers    int
	workerConfigs struct {
		batch          int
		maxIdleTime    time.Duration
		maxStoppedTime time.Duration
	}
	jobConfigs struct {
		maxAttempts              int
		retryDelay               time.Duration
		retryMaxDelay            time.Duration
		retryDelayStep           time.Duration
		maxExecuteTimePerAttempt time.Duration
		maxExecuteTimeTotal      time.Duration
	}
}

// DefaultWorkerPool creates WorkerPool instance with specific driver and default configs
// example:
//
//	wp := DefaultWorkerPool(driver.NewQueueDriver())
//	wp.Start()
func DefaultWorkerPool(driver IWorkerPoolDriver) *WorkerPool {
	wp := new(WorkerPool)
	// basic attributes
	wp.id = uuid.New()
	wp.status = WorkerPoolStatusStopped
	wp.createdAt = time.Now()
	wp.stoppedAt = time.Now()
	// worker container
	wp.workers = maps.NewHashMap[uuid.UUID, *Worker]()
	// stop signal channel
	wp.stopChannel = make(chan struct{})
	// configs
	wp.workerConfigs.batch = DefaultWorkerBatch
	wp.workerConfigs.maxIdleTime = DefaultWorkerMaxIdleTime
	wp.workerConfigs.maxStoppedTime = DefaultWorkerStoppedTimeout
	wp.jobConfigs.maxAttempts = DefaultJobMaxAttempts
	wp.jobConfigs.retryDelay = DefaultJobRetryDelay
	wp.jobConfigs.retryMaxDelay = DefaultJobRetryMaxDelay
	wp.jobConfigs.retryDelayStep = DefaultJobRetryDelayStep
	wp.jobConfigs.maxExecuteTimeTotal = DefaultJobMaxExecuteTimeTotal
	wp.jobConfigs.maxExecuteTimePerAttempt = DefaultJobMaxExecuteTimePerAttempt
	return wp
}

// NewWorkerPool creates a [WorkerPool] instance with specific driver and custom [Option]s
// example:
//
//	wp := NewWorkerPool(driver.NewQueueDriver(), MaxWorkers(10))
//	wp.Start()
func NewWorkerPool(driver IWorkerPoolDriver, options ...Option) *WorkerPool {
	wp := DefaultWorkerPool(driver)
	for _, option := range options {
		option(wp)
	}
	return wp
}

// NewWorkerWithConfigs creates a [WorkerPool] instance with specific driver and custom [WorkerPoolConfigs].
// In fact, it change configs to a list of [Option]s and calls [NewWorkerPool] to create WorkerPool
//
// example:
//
//	wp := NewWorkerWithConfigs(driver.NewQueueDriver(), &WorkerPoolConfigs{
//		MaxWorkers: 10
//	})
//	wp.Start()
func NewWorkerWithConfigs(driver IWorkerPoolDriver, configs *WorkerPoolConfigs) *WorkerPool {
	return NewWorkerPool(driver, configs.ToOptions()...)
}

// ID returns the unique id of the WorkerPool
func (wp *WorkerPool) ID() uuid.UUID {
	return wp.id
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
func (p *WorkerPool) Dispatch(job IJob) bool {
	if p.IsStopped() {
		return false
	}
	ok := p.driver.Enqueue(job)
	p.spawnWorkers()
	return ok
}

// Start starts the workerpool
func (p *WorkerPool) Start() {
	p.start()
	select {
	case <-p.stopChannel:
		p.workers.Range(func(entry *maps.Entry[uuid.UUID, *Worker]) bool {
			p.stopWorker(entry.Value)
			return true
		})
		p.stop()
	default:
		p.spawnWorkers()
		go p.watch()
	}
}

// Stop stops the worker pool and all the workers
func (p *WorkerPool) Stop() {
	p.stopChannel <- struct{}{}
}

// Release releases and removes the workerpool from the [WorkerPoolManager]
func (p *WorkerPool) Release() {
	if p.IsRunning() {
		p.Stop()
	}
	close(p.stopChannel)
}

func (wp *WorkerPool) shouldStopWorker(w *Worker) bool {
	return w.IsIdle() && time.Since(w.idledAt) >= wp.workerConfigs.maxIdleTime
}

func (wp *WorkerPool) shouldReleaseWorker(w *Worker) bool {
	return w.IsStopped() && time.Since(w.stoppedAt) >= wp.workerConfigs.maxStoppedTime
}

func (p *WorkerPool) stopWorker(w *Worker) {
	w.Stop()
}

func (p *WorkerPool) releaseWorker(w *Worker) {
	w.release()
	p.workers.Remove(w.id)
}

func (p *WorkerPool) watch() {
	for {
		wg := sync.WaitGroup{}
		workers := p.workers.Values()
		for _, w := range workers {
			wg.Add(1)
			go func(w *Worker) {
				defer wg.Done()
				if p.shouldStopWorker(w) {
					p.stopWorker(w)
				} else if p.shouldReleaseWorker(w) {
					p.releaseWorker(w)
				}
			}(w)
		}
		wg.Wait()
	}
}

func (wp *WorkerPool) spawnWorkers() {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	if wp.workers.Count() >= wp.maxWorkers {
		return
	}
	pendingCount := wp.driver.PendingCount()
	if pendingCount == 0 {
		return
	}
	nc := wp.maxWorkers / wp.workerConfigs.batch
	if nc == 0 {
		nc = wp.maxWorkers
	}
	if nc > int(pendingCount) {
		nc = int(pendingCount)
	}
	c := 0
	// awake sleeping workers
	wp.workers.Range(func(entry *maps.Entry[uuid.UUID, *Worker]) bool {
		if entry.Value.IsStopped() {
			go entry.Value.Start()
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
func (p *WorkerPool) Workers() []*Worker {
	return p.workers.Values()
}

// Progress returns active progress
func (p *WorkerPool) Progress() *Progress {
	return &Progress{
		Total:     p.driver.Total(),
		Pending:   p.driver.PendingCount(),
		Executing: p.driver.ExecutingCount(),
		Completed: p.driver.CompletedCount(),
		Success:   p.driver.SuccessCount(),
		Failed:    p.driver.FailedCount(),
	}
}
