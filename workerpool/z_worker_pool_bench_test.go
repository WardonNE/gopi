package workerpool

import (
	"testing"

	"github.com/wardonne/gopi/workerpool/driver"
)

func BenchmarkWorkerPool_Dispatch(b *testing.B) {
	wp := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(10))
	wp.Start()
	for i := 0; i < b.N; i++ {
		wp.Dispatch(&testjob{callback: func() error {
			return nil
		}})
	}
	for wp.driver.Count() > 0 {

	}
}
