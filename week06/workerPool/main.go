package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

const taskNum = 20
const workerNum = 5

func worker(jobs <-chan int, results chan<- int, id int) {
	for task := range jobs {
		delay := time.Duration(100+rand.IntN(301)) * time.Millisecond
		fmt.Printf("worker %d processing %d\n", id, task)
		time.Sleep(delay)
		results <- task
		fmt.Printf("worker %d finished %d\n", id, task)
	}
}

func main() {
	jobs := make(chan int, taskNum)
	results := make(chan int, taskNum)
	var wg sync.WaitGroup

	for i := 1; i <= taskNum; i++ {
		jobs <- i
	}
	close(jobs)

	for id := 1; id <= workerNum; id++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(jobs, results, id)
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for task := range results {
		fmt.Println(task)
	}
}
