package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	wg sync.WaitGroup
	x  int64
	lg sync.Mutex
)

// 普通操作
func add() {
	x += 1
	wg.Done()
}

// 互斥锁
func lockAdd() {
	lg.Lock()
	x += 1
	lg.Unlock()
	wg.Done()
}

// 原子操作相加
func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	start := time.Now()

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		//go add() // 400ms
		//go lockAdd() // 407ms
		go atomicAdd() // 385ms
	}

	end := time.Now()

	fmt.Println(end.Sub(start))
	fmt.Println(x)
	wg.Wait()
}
