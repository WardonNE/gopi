package workerpool

import "github.com/wardonne/gopi/support/maps"

type WorkerPoolManager struct {
	pools *maps.SyncHashMap[string, IWorkerPool]
}

// RegisterWorkerPool registers an existing worker pool
func (wpm *WorkerPoolManager) RegisterWorkerPool(name string, workerPool IWorkerPool) bool {
	if wpm.pools.ContainsKey(name) {
		return false
	}
	wpm.pools.Set(name, workerPool)
	return true
}

// RegisterWorkerPools list all registered pools
func (wpm *WorkerPoolManager) RegisterWorkerPools() *maps.SyncHashMap[string, IWorkerPool] {
	return wpm.pools
}

// GetWorkerPool returns Worker pool by the specific name
func (wpm *WorkerPoolManager) GetWorkerPool(name string) IWorkerPool {
	return wpm.pools.Get(name)
}

// CreateWorkerPool creates a new worker pool with max worker count and registers it with the specific name
func (wpm *WorkerPoolManager) CreateWorkerPool(name string, cap int, driver IWorkerPoolDriver) (workerPool IWorkerPool, isNew bool) {
	if wpm.pools.ContainsKey(name) {
		return wpm.pools.Get(name), false
	}
	workerPool = NewWorkerPool(cap, driver)
	wpm.pools.Set(name, workerPool)
	return workerPool, true
}

// StartWorkerPool starts the specific worker pool
func (wpm *WorkerPoolManager) StartWorkerPool(name string) {
	if wpm.pools.ContainsKey(name) {
		workerPool := wpm.pools.Get(name)
		if !workerPool.Running() {
			workerPool.Start()
		}
	}
}

// StopWorkerPool stops the specific worker pool
func (wpm *WorkerPoolManager) StopWorkerPool(name string) {
	if wpm.pools.ContainsKey(name) {
		workerPool := wpm.pools.Get(name)
		if workerPool.Running() {
			workerPool.Stop()
		}
	}
}

// ReleaseWorkerPool releases the specific worker pool
func (wpm *WorkerPoolManager) ReleaseWorkerPool(name string) {
	if wpm.pools.ContainsKey(name) {
		workerPool := wpm.pools.Get(name)
		if workerPool.Running() {
			workerPool.Release()
		}
		wpm.pools.Remove(name)
	}
}

// StartAll starts all worker pools
func (wpm *WorkerPoolManager) StartAll() {
	wpm.pools.Range(func(entry *maps.Entry[string, IWorkerPool]) bool {
		if !entry.Value.Running() {
			go entry.Value.Start()
		}
		return true
	})
}

// StopAll stops all worker pools
func (wpm *WorkerPoolManager) StopAll() {
	wpm.pools.Range(func(entry *maps.Entry[string, IWorkerPool]) bool {
		if entry.Value.Running() {
			go entry.Value.Stop()
		}
		return true
	})
}

// ReleaseAll releases all worker pools
func (wpm *WorkerPoolManager) ReleaseAll() {
	wpm.pools.Range(func(entry *maps.Entry[string, IWorkerPool]) bool {
		go entry.Value.Release()
		return true
	})
	wpm.pools.Clear()
}
