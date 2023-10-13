package workerpool

import (
	"github.com/wardonne/gopi/support/maps"
	"github.com/wardonne/gopi/workerpool/driver"
)

type WorkerPoolManager struct {
	pools *maps.SyncHashMap[string, *WorkerPool]
}

func NewWorkerPoolManager() *WorkerPoolManager {
	manager := new(WorkerPoolManager)
	manager.pools = maps.NewSyncHashMap[string, *WorkerPool]()
	return manager
}

// List lists all registered worker pools
func (wpm *WorkerPoolManager) List() map[string]*WorkerPool {
	return wpm.pools.ToMap()
}

// Get returns Worker pool by the specific name
func (wpm *WorkerPoolManager) Get(name string) *WorkerPool {
	return wpm.pools.Get(name)
}

// Create creates a new worker pool with max worker count and registers it with the specific name
func (wpm *WorkerPoolManager) Create(name string, driver driver.DriverInterface, options ...Option) (workerPool *WorkerPool, isNew bool) {
	if wpm.pools.ContainsKey(name) {
		return wpm.pools.Get(name), false
	}
	workerPool = NewWorkerPool(driver, options...)
	workerPool.name = name
	wpm.pools.Set(name, workerPool)
	return workerPool, true
}

// Add registers an existing worker pool
func (wpm *WorkerPoolManager) Add(name string, workerPool *WorkerPool) (bool, error) {
	if wpm.pools.ContainsKey(name) {
		return false, ErrWorkerPoolNameExists
	}
	if wpm.pools.ContainsValue(func(value *WorkerPool) bool {
		return workerPool.id == value.id
	}) {
		return false, ErrWorkerPoolInstanceExists
	}
	wpm.pools.Set(name, workerPool)
	workerPool.name = name
	return true, nil
}

// Start starts the specific worker pool
func (wpm *WorkerPoolManager) Start(name string) {
	if wpm.pools.ContainsKey(name) {
		workerPool := wpm.pools.Get(name)
		if workerPool.IsStopped() {
			workerPool.Start()
		}
	}
}

// Stop stops the specific worker pool
func (wpm *WorkerPoolManager) Stop(name string) {
	if wpm.pools.ContainsKey(name) {
		workerPool := wpm.pools.Get(name)
		if workerPool.IsRunning() {
			workerPool.Stop()
		}
	}
}

// Release releases the specific worker pool
func (wpm *WorkerPoolManager) Release(name string) {
	if wpm.pools.ContainsKey(name) {
		workerPool := wpm.pools.Get(name)
		if workerPool.IsRunning() {
			workerPool.Release()
		}
		wpm.pools.Remove(name)
	}
}

// StartAll starts all worker pools
func (wpm *WorkerPoolManager) StartAll() {
	wpm.pools.Range(func(entry *maps.Entry[string, *WorkerPool]) bool {
		if entry.Value.IsStopped() {
			go entry.Value.Start()
		}
		return true
	})
}

// StopAll stops all worker pools
func (wpm *WorkerPoolManager) StopAll() {
	wpm.pools.Range(func(entry *maps.Entry[string, *WorkerPool]) bool {
		if entry.Value.IsRunning() {
			go entry.Value.Stop()
		}
		return true
	})
}

// ReleaseAll releases all worker pools
func (wpm *WorkerPoolManager) ReleaseAll() {
	wpm.pools.Range(func(entry *maps.Entry[string, *WorkerPool]) bool {
		go entry.Value.Release()
		return true
	})
	wpm.pools.Clear()
}
