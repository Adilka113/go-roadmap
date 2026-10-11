package main

import "fmt"

func receive(ch <-chan string) {
	fmt.Println("Получено:", <-ch)
}

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Ping"
	}()
	receive(ch)
}
