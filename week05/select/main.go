package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go func() {
		ch1 <- "fast"
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- "slow"
	}()

	select {
	case v := <-ch1:
		fmt.Println("Первым ответил:", v)
	case v := <-ch2:
		fmt.Println("Первым ответил:", v)
	}
}
