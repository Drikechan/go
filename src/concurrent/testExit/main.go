package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		defer fmt.Println("A")
		func() {
			defer fmt.Println("B")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C")
			fmt.Println("b")
		}()

		fmt.Println("a")
	}()
	for {

	}
}
