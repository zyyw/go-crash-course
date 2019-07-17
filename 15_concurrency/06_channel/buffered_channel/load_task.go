package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	// create a buffered channel to manage the task load
	tasks := make(chan string, taskLoad)

	// launch goroutines to handle the work
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// add a bunch of work to get done
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}

	// close the channel so the goroutines will quit when all tasks are done by works
	close(tasks)

	// wait for all the work to get done
	wg.Wait()
}

func worker(tasks chan string, id int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			// this means the channel is empty and closed
			fmt.Printf("Worker: %d shut down\n", id)
			return
		}

		// display we start the work
		fmt.Printf("Worker %d: started %s\n", id, task)

		// randomly wait to simulate work time
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))

		// display we finished the work
		fmt.Printf("Worker %d: completed %s\n", id, task)
	}
}
