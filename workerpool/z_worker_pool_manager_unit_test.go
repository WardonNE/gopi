package workerpool

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wardonne/gopi/workerpool/driver"
)

func TestWorkerPoolManager_Create(t *testing.T) {
	wpm := NewWorkerPoolManager()
	wp, isNew := wpm.Create("wp", driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	assert.True(t, isNew)
	assert.Equal(t, "wp", wp.Name())
	assert.Equal(t, wpm.Get("wp"), wp)

	t.Run("Create-RepeatName", func(t *testing.T) {
		wp1, isNew := wpm.Create("wp", driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
		assert.False(t, isNew)
		assert.Equal(t, "wp", wp1.Name())
		assert.Equal(t, wp, wp1)
	})
}

func TestWorkerPoolManager_Add(t *testing.T) {
	wpm := NewWorkerPoolManager()

	wp := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	success, err := wpm.Add("wp", wp)
	assert.True(t, success)
	assert.Nil(t, err)
	assert.Equal(t, "wp", wp.Name())
	assert.Equal(t, wpm.Get("wp"), wp)

	t.Run("Add-RepeatName", func(t *testing.T) {
		success, err := wpm.Add("wp", wp)
		assert.False(t, success)
		assert.Equal(t, ErrWorkerPoolNameExists, err)
	})

	t.Run("Add-RepeatWorkerPool", func(t *testing.T) {
		success, err := wpm.Add("wp1", wp)
		assert.False(t, success)
		assert.Equal(t, ErrWorkerPoolInstanceExists, err)
	})
}

func TestWorkerPoolManager_List(t *testing.T) {
	wpm := NewWorkerPoolManager()

	wp0, isNew := wpm.Create("wp0", driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	assert.True(t, isNew)
	assert.Equal(t, "wp0", wp0.Name())
	assert.Equal(t, wpm.Get("wp0"), wp0)

	wp1 := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	success, err := wpm.Add("wp1", wp1)
	assert.True(t, success)
	assert.Nil(t, err)

	wp2 := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	success, err = wpm.Add("wp2", wp2)
	assert.True(t, success)
	assert.Nil(t, err)

	assert.Equal(t, map[string]*WorkerPool{
		"wp0": wp0,
		"wp1": wp1,
		"wp2": wp2,
	}, wpm.List())
}

func TestWorkerPoolManager_Start(t *testing.T) {
	wpm := NewWorkerPoolManager()

	wp0, _ := wpm.Create("wp0", driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))

	wp1 := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	wpm.Add("wp1", wp1)

	wp2 := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	wpm.Add("wp2", wp2)

	wpm.Start("wp1")

	assert.True(t, wp0.IsStopped())
	assert.True(t, wp1.IsRunning())
	assert.True(t, wp2.IsStopped())
}

func TestWorkerPoolManager_StartAll(t *testing.T) {
	wpm := NewWorkerPoolManager()

	wp0, _ := wpm.Create("wp0", driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))

	wp1 := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	wpm.Add("wp1", wp1)

	wp2 := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	wpm.Add("wp2", wp2)

	wpm.StartAll()
	time.Sleep(time.Second)
	assert.True(t, wp0.IsRunning())
	assert.True(t, wp1.IsRunning())
	assert.True(t, wp2.IsRunning())
}

func TestWorkerPoolManager_Stop(t *testing.T) {
	wpm := NewWorkerPoolManager()

	wp0, _ := wpm.Create("wp0", driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))

	wp1 := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	wpm.Add("wp1", wp1)

	wp2 := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	wpm.Add("wp2", wp2)

	wpm.StartAll()
	time.Sleep(time.Second)
	assert.True(t, wp0.IsRunning())
	assert.True(t, wp1.IsRunning())
	assert.True(t, wp2.IsRunning())

	wpm.Stop("wp0")
	assert.True(t, wp0.IsStopped())
	assert.True(t, wp1.IsRunning())
	assert.True(t, wp2.IsRunning())
}

func TestWorkerPoolManager_StopAll(t *testing.T) {
	wpm := NewWorkerPoolManager()

	wp0, _ := wpm.Create("wp0", driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))

	wp1 := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	wpm.Add("wp1", wp1)

	wp2 := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	wpm.Add("wp2", wp2)

	wpm.StartAll()
	time.Sleep(time.Second)
	assert.True(t, wp0.IsRunning())
	assert.True(t, wp1.IsRunning())
	assert.True(t, wp2.IsRunning())

	wpm.StopAll()
	time.Sleep(time.Second)
	assert.True(t, wp0.IsStopped())
	assert.True(t, wp1.IsStopped())
	assert.True(t, wp2.IsStopped())
}

func TestWorkerPoolManager_Release(t *testing.T) {
	wpm := NewWorkerPoolManager()

	wp0, _ := wpm.Create("wp0", driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))

	wp1 := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	wpm.Add("wp1", wp1)

	wp2 := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	wpm.Add("wp2", wp2)

	wpm.StartAll()
	time.Sleep(time.Second)
	assert.True(t, wp0.IsRunning())
	assert.True(t, wp1.IsRunning())
	assert.True(t, wp2.IsRunning())

	wpm.Release("wp0")
	assert.True(t, wp0.IsStopped())
	assert.True(t, wp1.IsRunning())
	assert.True(t, wp2.IsRunning())

	assert.Nil(t, wpm.Get("wp0"))
}

func TestWorkerPoolManager_ReleaseAll(t *testing.T) {
	wpm := NewWorkerPoolManager()

	wp0, _ := wpm.Create("wp0", driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))

	wp1 := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	wpm.Add("wp1", wp1)

	wp2 := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	wpm.Add("wp2", wp2)

	wpm.StartAll()
	time.Sleep(time.Second)
	assert.True(t, wp0.IsRunning())
	assert.True(t, wp1.IsRunning())
	assert.True(t, wp2.IsRunning())

	wpm.ReleaseAll()
	time.Sleep(time.Second)
	assert.True(t, wp0.IsStopped())
	assert.True(t, wp1.IsStopped())
	assert.True(t, wp2.IsStopped())

	assert.Equal(t, map[string]*WorkerPool{}, wpm.List())
}
