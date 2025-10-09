package concurrency

import (
	"fmt"
	"sync"
)

// implementation of worker pool from scratch
type WorkerJob struct {
	Id            uint64
	ExecutionUnit func()
}

// constructor for creating new worker job
func NewWorkerJob(id uint64, eu func()) *WorkerJob {
	return &WorkerJob{
		Id:            id,
		ExecutionUnit: eu,
	}
}

// simulate the scenario for 100 jobs and 10 workers
func Simulate() {
	nj := 100
	nw := 10
	jc := make(chan *WorkerJob, nj)
	var wg sync.WaitGroup

	// worker execution routine
	wf := func(jc <-chan *WorkerJob, wid int, wg *sync.WaitGroup) {
		defer wg.Done()
		for j := range jc {
			fmt.Printf("Executing job - %d  by worker - %d\n", j.Id, wid)
			fmt.Printf("Finished execution of job - %d by worker - %d\n", j.Id, wid)
		}
	}

	// spin the workers
	for i := 0; i < nw; i++ {
		wg.Add(1)
		go wf(jc, i+1, &wg)
	}

	// create jobs and populate them in channels
	for i := 0; i < nj; i++ {
		wj := NewWorkerJob(uint64(i+1), func() {
			// dummy function
			var mu sync.Mutex
			// handling the race condition...
			mu.Lock()
			fmt.Println("your business logic goes here!")
			mu.Unlock()

		})
		jc <- wj
	}

	close(jc)
	wg.Wait()
}
