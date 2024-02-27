package main

import "sync"

func drainChannel[T any](in chan T, n int) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			for data := range in {
				_ = data
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
