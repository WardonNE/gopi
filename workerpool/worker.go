package workerpool

type Worker struct {
	driver IWorkerPoolDriver
	done   chan struct{}
	stop   chan struct{}
}

func NewWorker(pool IWorkerPool) *Worker {
	worker := new(Worker)
	worker.driver = pool.Driver()
	worker.done = make(chan struct{})
	worker.stop = make(chan struct{})
	return worker
}

func (w *Worker) Start() {
	for {
		select {
		case <-w.stop:
			return
		default:
			if job, ok := w.driver.Dequeue(); ok {
				job.Handle()
			}
		}
	}
}

func (w *Worker) Stop() {
	w.stop <- struct{}{}
}
