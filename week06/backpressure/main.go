package main

import (
	"fmt"
	"sync"
	"time"
)

const consumerDelay = time.Millisecond * 200
const bufSize = 5
const prodCount = 20

func producer(ch chan<- int) {
	defer close(ch)
	for i := 1; i <= prodCount; i++ {
		select {
		case ch <- i:
			fmt.Println("produced:", i)
		default:
			fmt.Println("BUFFER FULL -> producer ждет")
			ch <- i
			fmt.Println("produced:", i)
		}
	}
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("consumed:", v)
		time.Sleep(consumerDelay)
	}
}

func main() {
	ch := make(chan int, bufSize)
	var wg sync.WaitGroup

	wg.Add(1)

	go producer(ch)
	go consumer(ch, &wg)

	wg.Wait()
}
