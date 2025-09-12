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
	// read from one channel and
}
