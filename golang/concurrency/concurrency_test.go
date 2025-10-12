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

func TestSelectUsage(t *testing.T) {
	SelectUsage()
}

func TestBasicExecutor(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go basicExecute(&wg)
	wg.Wait()
}
func TestSimulate(t *testing.T) {
	Simulate()
}

func TestSimulateProducerConsumer(t *testing.T) {
	SimulateProducerConsumer()
}

func TestThreadSafeSingleton(t *testing.T) {
	ThreadSafeSingleton()
}
