package main

import (
	"fmt"  // пакет используется для проверки выполнения условия задачи, не удаляйте его
	"time" // пакет используется для проверки выполнения условия задачи, не удаляйте его
    "sync"
)

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	var (
		slice1   = make([]int, n)
		slice2   = make([]int, n)
		wg sync.WaitGroup
		mu sync.Mutex
	)
	wg.Add(n * 2)
	go func() {
		for i := 0; i < n; i++ {
			mu.Lock()
			slice1[i] = <-in1
			mu.Unlock()
			wg.Done()
		}
	}()
	go func() {
		for i := 0; i < n; i++ {
			mu.Lock()
			slice2[i] = <-in2
			mu.Unlock()
			wg.Done()
		}
	}()
	go func() {
		for i, v := range slice1 {
			slice1[i] = fn(v)
		}
	}()
	go func() {
		for i, v := range slice2 {
			slice2[i] = fn(v)
		}
	}()
	go func() {
		defer close(out)
		wg.Wait()
		for i := 0; i < n; i++ {
			out <- slice1[i] + slice2[i]
		}
	}()
}