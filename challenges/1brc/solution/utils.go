package main

import (
	"sync"
)

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

type Pool struct {
	data [][]byte
}

func NewPool(initialSize int) *Pool {
	return &Pool{data: make([][]byte, initialSize)}
}

func (p *Pool) Get() (last []byte) {
	if len(p.data) == 0 {
		return make([]byte, ioBufferSize)
	}
	last = p.data[len(p.data)-1]
	p.data = p.data[:len(p.data)-1]
	return last
}

func (p *Pool) Put(x []byte) {
	p.data = append(p.data, x[:0])
}
