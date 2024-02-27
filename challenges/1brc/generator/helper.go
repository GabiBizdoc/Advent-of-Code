package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const itoaSize = 1000

type CustomItoa struct {
	m            atomic.Bool
	positiveItoa []string
}

func NewCustomItoa() *CustomItoa {
	positiveItoa := make([]string, itoaSize*2)
	c := &CustomItoa{positiveItoa: positiveItoa}
	c.init()
	return c
}

func (c *CustomItoa) init() {
	go func() {
		for i := 0; i < itoaSize; i++ {
			c.positiveItoa[i] = strconv.Itoa(i)
		}
		for i := itoaSize; i < itoaSize*2; i++ {
			c.positiveItoa[i] = "-" + strconv.Itoa(i)
		}
		c.m.Store(true)
	}()
}

func (c *CustomItoa) Parse(n int) string {
	if c.m.Load() == false {
		for c.m.Load() == false {
			time.Sleep(time.Millisecond)
		}
	}

	if n < 0 {
		n = itoaSize - n
	}

	if n >= itoaSize*2 {
		panic(fmt.Sprintf("invalid temperature %d", n))
	}
	return c.positiveItoa[n]
}

// SlicePool not used
type SlicePool struct {
	slicePool sync.Pool
}

func (s *SlicePool) Get() []byte {
	return s.slicePool.Get().([]byte)
}

func (s *SlicePool) Put(b []byte) {
	b = b[0:0]
	s.slicePool.Put(b)
}

func NewSlicePool() *SlicePool {
	return &SlicePool{slicePool: sync.Pool{
		New: func() interface{} {
			return make([]byte, 0, 10000)
		},
	}}
}
