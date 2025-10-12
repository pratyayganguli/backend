package concurrency

import (
	"fmt"
	"strconv"
	"sync"
)

// requirement: build a producer consumer design using concurrency. The producer will send the events and the consumer will accept the messages and process them

type Event struct {
	ExecutionUnit func(string) string
}

func NewEvent(message string) *Event {
	f := func(string) string {
		return fmt.Sprintf("Sending message %s", message)
	}
	return &Event{
		ExecutionUnit: f,
	}
}

type Producer struct {
	Channel chan Event
}

func (p *Producer) GetChannel() chan Event {
	if p.Channel != nil {
		panic("event channel is nil!")
	} else {
		return p.Channel
	}
}

type Consumer struct {
	Id      string
	Channel chan Event
}

func SimulateProducerConsumer() {
	n := 10
	ec := make(chan *Event, n)
	var wg sync.WaitGroup

	w := func(wg *sync.WaitGroup, messages <-chan *Event, id int) {
		defer wg.Done()
		for m := range messages {
			m.ExecutionUnit("abc")
			fmt.Printf("Processing message by consumer %d", id)
		}
	}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go w(&wg, ec, i)
	}

	// Producer is sending messages in the event channel
	for i := 0; i < n; i++ {
		e := NewEvent(strconv.Itoa(i + 1))
		ec <- e
	}
	close(ec)
	wg.Wait()
}
