package main

import "fmt"

func squares(n int, ch chan<- int) {
	defer close(ch)
	for i := 1; i <= n; i++ {
		ch <- i * i
	}
}

func main() {
	const ln = 5
	var sum int
	ch := make(chan int, ln)

	go squares(ln, ch)

	for v := range ch {
		sum += v
	}
	fmt.Println(sum)
}
