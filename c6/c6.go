package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu      sync.Mutex // mutex protect queue
	cond    = sync.NewCond(&mu)
	queue   []int
	maxSize = 3
)

func main() {
	fmt.Println("module: c6")
	fmt.Println("version: 0.0.1")

	go produce(1)
	go consume(1)

	time.Sleep(time.Second * 5)
}

func produce(id int) {

	for i := 0; i < 10; i++ {

		mu.Lock()
		for len(queue) == maxSize {
			fmt.Printf("Producer %d is waiting, queue full\n", id)
			cond.Wait()
		}

		queue = append(queue, i)
		fmt.Printf("Producer %d produced %d -> queue: %v\n", id, i, queue)

		cond.Signal() // wake up consumer
		mu.Unlock()
		time.Sleep(time.Microsecond * 250) // produce faster
	}

}

func consume(id int) {

	for {

		mu.Lock()
		for len(queue) == 0 {
			fmt.Printf("Consumer %d is waiting, queue empty\n", id)
			cond.Wait()
		}

		item := queue[0]
		queue = queue[1:]

		fmt.Printf("Consumer %d consumed %d -> queue: %v\n", id, item, queue)

		cond.Signal() // wake up producer
		mu.Unlock()
		time.Sleep(time.Second * 1) // consume slower
	}

}
