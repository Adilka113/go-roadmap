package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const delay = time.Second
const tickInterval = 200 * time.Millisecond

func worker(ctx context.Context) {
	ticker := time.NewTicker(tickInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context canceled, stopping")
			return
		case <-ticker.C:
			fmt.Println("working...")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		worker(ctx)
	}()
	time.Sleep(delay)
	cancel()
	wg.Wait()
}
