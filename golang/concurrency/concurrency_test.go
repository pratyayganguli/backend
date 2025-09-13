package concurrency

import (
	"fmt"
	"sync"
	"testing"
)

func basicExecute(wg *sync.WaitGroup) {
	// Done is bound to be invoked even though the execution might fail
	defer wg.Done()
	fmt.Println("Task executed!")
}

func TestBasicExecutor(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go basicExecute(&wg)
	wg.Wait()
}

// channels in golang are custom types though which go routines tranfer data with each other
func TestOddEvenPrint(t *testing.T) {

	// unbufferred channels created with make(), they do not have a fixed size
	// helps to provide synchronous communication between sender and receiver

	oddChan := make(chan int)
	evenChan := make(chan int)
	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for num := range oddChan {
			fmt.Printf("odd number %d\n", num)
		}
	}()

	go func() {
		defer wg.Done()
		for num := range evenChan {
			fmt.Printf("even number %d\n", num)
		}
	}()

	for i := range input {
		if input[i]%2 == 0 {
			evenChan <- input[i]
		} else {
			oddChan <- input[i]
		}
	}

	close(oddChan)
	close(evenChan)
	wg.Wait()
}

func TestOddThenEven(t *testing.T) {
	oddChan := make(chan int)
	evenChan := make(chan int)
	done := make(chan struct{})
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var wg sync.WaitGroup

	// odd receiver routine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for oddNum := range oddChan {
			fmt.Printf("Odd number: %d\n", oddNum)
			evenChan <- oddNum + 1
		}
		close(evenChan)
	}()

	// even receiver routine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for evenNum := range evenChan {
			fmt.Printf("Even number: %d\n", evenNum)
		}
		close(done)
	}()

	go func() {
		for _, num := range nums {
			if num%2 != 0 {
				oddChan <- num
			}
		}
		close(oddChan)
	}()

	<-done
	wg.Wait()
}

// the job channels hold the tasks which needs to be executed
type Priority string

const MEDIUM_PRIORITY Priority = "medium"
const CRITICAL_PRIORITY Priority = "critical"

type Job struct {
	Id       int
	Priority Priority
}

func TestWorkerPool(t *testing.T) {
	const MAX_WORKERS = 10
	const MAX_JOBS = 20
	var wg sync.WaitGroup

	jobsChan := make(chan *Job, MAX_JOBS)

	// worker function for executing all the jobs
	w := func(id int, jobs <-chan *Job, wg *sync.WaitGroup) {
		defer wg.Done()
		for job := range jobs {
			fmt.Printf("Worker %d executing - job %d on priority %s\n", id, job.Id, job.Priority)
		}
	}

	// spawn the worker go routines
	for i := 1; i <= MAX_WORKERS; i++ {
		wg.Add(1)
		go w(i, jobsChan, &wg)
	}

	// send the jobs to the job channel
	for i := 0; i < MAX_JOBS; i++ {
		j := &Job{
			Id:       i,
			Priority: MEDIUM_PRIORITY,
		}
		jobsChan <- j
	}

	close(jobsChan)
	wg.Wait()
}
