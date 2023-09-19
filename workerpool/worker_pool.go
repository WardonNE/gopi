package workerpool

import (
	"github.com/wardonne/gopi/support/collection/list"
)

var _ IWorkerPool = (*WorkerPool)(nil)

type WorkerPool struct {
	driver  IWorkerPoolDriver
	workers *list.ArrayList[*Worker]
	cap     int
	running bool
	stop    chan struct{}
}

func NewWorkerPool(cap int, driver IWorkerPoolDriver) *WorkerPool {
	p := new(WorkerPool)
	p.driver = driver
	p.workers = list.NewArrayList[*Worker]()
	p.cap = cap
	p.stop = make(chan struct{})
	return p
}

func (p *WorkerPool) Driver() IWorkerPoolDriver {
	return p.driver
}

func (p *WorkerPool) Dispatch(job Job) bool {
	return p.driver.Enqueue(job)
}

func (p *WorkerPool) Running() bool {
	return p.running
}

func (p *WorkerPool) Start() {
	p.running = true
	if p.workers.IsEmpty() {
		for i := 0; i < p.cap; i++ {
			worker := NewWorker(p)
			p.workers.Add(worker)
		}
	}
	p.workers.Range(func(worker *Worker) bool {
		go worker.Start()
		return true
	})
	<-p.stop
	p.workers.Range(func(worker *Worker) bool {
		go worker.Stop()
		return true
	})
	p.running = false
}

func (p *WorkerPool) Stop() {
	p.stop <- struct{}{}
}

func (p *WorkerPool) Release() {
	if p.running {
		p.Stop()
	}
	close(p.stop)
}
