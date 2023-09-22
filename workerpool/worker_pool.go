package workerpool

import (
	"context"
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
	name      string
	status    WorkerPoolStatus
	createdAt time.Time
	startAt   time.Time
	stoppedAt time.Time

	mu          sync.Mutex
	workers     *maps.HashMap[uuid.UUID, *Worker]
	stopChannel chan struct{}

	driver     IWorkerPoolDriver
	maxWorkers int
	// worker configs
	workerConfigs struct {
		batch          int
		maxIdleTime    time.Duration
		maxStoppedTime time.Duration
	}
	// job configs
	jobConfigs struct {
		maxAttempts              int
		retryDelay               time.Duration
		retryMaxDelay            time.Duration
		retryDelayStep           time.Duration
		maxExecuteTimePerAttempt time.Duration
		maxExecuteTimeTotal      time.Duration
	}
	// listeners
	progressListener func(*Progress)
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
func (wp *WorkerPool) Dispatch(job IJob) bool {
	if wp.IsStopped() {
		return false
	}
	ok := wp.driver.Enqueue(job)
	if wp.progressListener != nil {
		wp.progressListener(wp.Progress())
	}
	wp.spawnWorkers()
	return ok
}

// Start starts the workerpool
func (wp *WorkerPool) Start() {
	wp.start()
	watchCtx, watchCancel := context.WithCancel(context.Background())
	defer watchCancel()
	listenerCtx, listenerCancel := context.WithCancel(context.Background())
	defer listenerCancel()
	select {
	case <-wp.stopChannel:
		wp.workers.Range(func(entry *maps.Entry[uuid.UUID, *Worker]) bool {
			wp.stopWorker(entry.Value)
			return true
		})
		wp.stop()
	default:
		wp.spawnWorkers()
		go wp.watch(watchCtx)
		go wp.notify(listenerCtx)
	}
}

// Stop stops the worker pool and all the workers
func (wp *WorkerPool) Stop() {
	wp.stopChannel <- struct{}{}
}

// Release releases and removes the workerpool from the [WorkerPoolManager]
func (wp *WorkerPool) Release() {
	if wp.IsRunning() {
		wp.Stop()
	}
	close(wp.stopChannel)
}

func (wp *WorkerPool) shouldStopWorker(w *Worker) bool {
	return w.IsIdle() && time.Since(w.idledAt) >= wp.workerConfigs.maxIdleTime
}

func (wp *WorkerPool) shouldReleaseWorker(w *Worker) bool {
	return w.IsStopped() && time.Since(w.stoppedAt) >= wp.workerConfigs.maxStoppedTime
}

func (wp *WorkerPool) stopWorker(w *Worker) {
	w.Stop()
}

func (wp *WorkerPool) releaseWorker(w *Worker) {
	w.release()
	wp.workers.Remove(w.id)
}

func (wp *WorkerPool) notify(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if wp.progressListener != nil {
				wp.progressListener(wp.Progress())
			}
		}
		time.Sleep(time.Second)
	}
}

func (wp *WorkerPool) watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			wg := sync.WaitGroup{}
			workers := wp.workers.Values()
			for _, w := range workers {
				wg.Add(1)
				go func(w *Worker) {
					defer wg.Done()
					if wp.shouldStopWorker(w) {
						wp.stopWorker(w)
					} else if wp.shouldReleaseWorker(w) {
						wp.releaseWorker(w)
					}
				}(w)
			}
			wg.Wait()
		}
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
func (wp *WorkerPool) Workers() []*Worker {
	return wp.workers.Values()
}

// Progress returns active progress
func (wp *WorkerPool) Progress() *Progress {
	return &Progress{
		Total:     wp.driver.Total(),
		Pending:   wp.driver.PendingCount(),
		Executing: wp.driver.ExecutingCount(),
		Completed: wp.driver.CompletedCount(),
		Success:   wp.driver.SuccessCount(),
		Failed:    wp.driver.FailedCount(),
	}
}
