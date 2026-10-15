package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Go(func() { //  в go 1.25 появилась замена Add/Done
			fmt.Println("Работаю...")
			time.Sleep(time.Second)
		})

		// Классический вариант до 1.25
		// wg.Add(1)
		// go func () {
		// 	defer wg.Done()
		// 	fmt.Println("Работаю...")
		// 	time.Sleep(time.Second)
		// }()
	}

	wg.Wait()
}
