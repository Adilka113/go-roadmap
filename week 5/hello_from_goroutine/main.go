package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 3; i++ {
		go fmt.Printf("Hello from goroutine #%d\n", i)
	}
	time.Sleep(time.Second)
}
