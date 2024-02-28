package workerpool

import (
	"github.com/wardonne/gopi/support/maps"
	"github.com/wardonne/gopi/workerpool/driver"
)

// Manager workerpool manager
type Manager struct {
	pools *maps.SyncHashMap[string, *WorkerPool]
}

// NewManager creates a new workerpool manager
func NewManager() *Manager {
	manager := new(Manager)
	manager.pools = maps.NewSyncHashMap[string, *WorkerPool]()
	return manager
}

// List lists all registered worker pools
func (wpm *Manager) List() map[string]*WorkerPool {
	return wpm.pools.ToMap()
}

// Get returns Worker pool by the specific name
func (wpm *Manager) Get(name string) *WorkerPool {
	return wpm.pools.Get(name)
}

// Create creates a new worker pool with max worker count and registers it with the specific name
func (wpm *Manager) Create(name string, driver driver.IDriver, options ...Option) (workerPool *WorkerPool, isNew bool) {
	if wpm.pools.ContainsKey(name) {
		return wpm.pools.Get(name), false
	}
	workerPool = NewWorkerPool(driver, options...)
	workerPool.name = name
	wpm.pools.Set(name, workerPool)
	return workerPool, true
}

// Add registers an existing worker pool
func (wpm *Manager) Add(name string, workerPool *WorkerPool) (bool, error) {
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
func (wpm *Manager) Start(name string) {
	if wpm.pools.ContainsKey(name) {
		workerPool := wpm.pools.Get(name)
		if workerPool.IsStopped() {
			workerPool.Start()
		}
	}
}

// Stop stops the specific worker pool
func (wpm *Manager) Stop(name string) {
	if wpm.pools.ContainsKey(name) {
		workerPool := wpm.pools.Get(name)
		if workerPool.IsRunning() {
			workerPool.Stop()
		}
	}
}

// Release releases the specific worker pool
func (wpm *Manager) Release(name string) {
	if wpm.pools.ContainsKey(name) {
		workerPool := wpm.pools.Get(name)
		if workerPool.IsRunning() {
			workerPool.Release()
		}
		wpm.pools.Remove(name)
	}
}

// StartAll starts all worker pools
func (wpm *Manager) StartAll() {
	wpm.pools.Range(func(entry *maps.Entry[string, *WorkerPool]) bool {
		if entry.Value.IsStopped() {
			go entry.Value.Start()
		}
		return true
	})
}

// StopAll stops all worker pools
func (wpm *Manager) StopAll() {
	wpm.pools.Range(func(entry *maps.Entry[string, *WorkerPool]) bool {
		if entry.Value.IsRunning() {
			go entry.Value.Stop()
		}
		return true
	})
}

// ReleaseAll releases all worker pools
func (wpm *Manager) ReleaseAll() {
	wpm.pools.Range(func(entry *maps.Entry[string, *WorkerPool]) bool {
		go entry.Value.Release()
		return true
	})
	wpm.pools.Clear()
}
