package main

import (
	"fmt"
	"sync"
)

func hello(x int) {
	defer wg.Done()
	fmt.Println(x)
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}
