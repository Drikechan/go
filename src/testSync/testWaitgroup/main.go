package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func test() {
	defer wg.Done()
	fmt.Println("1111")
}

func main() {
	wg.Add(1)
	go test()
	wg.Wait()
	fmt.Println("2222")

}
