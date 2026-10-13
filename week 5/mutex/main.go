package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Store struct {
	mu sync.RWMutex
	m  map[string]string
}

func (s *Store) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m[key]
}

func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func main() {
	var wg sync.WaitGroup
	m := make(map[string]string)
	const setDelay = time.Millisecond * 10
	const getDelay = time.Millisecond * 5

	s := Store{m: m}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			s.Set("config", "v"+strconv.Itoa(i))
			time.Sleep(setDelay)
		}
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				fmt.Printf("Читатель #%d видит: %s\n", i, s.Get("config"))
				time.Sleep(getDelay)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Итог:", s.Get("config"))
}
