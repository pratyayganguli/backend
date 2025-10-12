package concurrency

import (
	"fmt"
	"time"
)

func SelectUsage() {
	ch := make(chan int)
	cch := make(chan int)

	// producer go routine
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
			time.Sleep(200 * time.Millisecond)
		}
		close(ch)
	}()

	go func() {
		// avoid using infinite loop
		for {
			select {
			case val, ok := <-ch:
				if !ok {
					close(cch)
					return
				}

				fmt.Println("data received on channel - ", val)
				// sending data to copy channel
				cch <- val
			}
		}
	}()

	for val := range cch {
		fmt.Println("data received on copy channel", val)
	}

}
