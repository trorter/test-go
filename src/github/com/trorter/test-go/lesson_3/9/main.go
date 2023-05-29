package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type atomicCounter struct {
	value int64
}

func (c *atomicCounter) Add(add int64) {
	atomic.AddInt64(&c.value, add)
}

func (c *atomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	counter := atomicCounter{}

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(no int) {
			for i := 0; i < 10_000; i++ {
				counter.Add(1)
				runtime.Gosched()
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(counter.Value())
}
