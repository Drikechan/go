package main

import (
	"fmt"
	"sync"
	"time"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 10000000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	fmt.Println(time.Now())
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
	fmt.Println(time.Now())
	fmt.Println("end")
}
