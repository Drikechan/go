package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		i := 0
		for {
			i++
			fmt.Println("goroutine里面的i", i)
			time.Sleep(time.Second)
		}
	}()

	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Second)
		if i == 2 {
			fmt.Println("main task end")
			break
		}
	}

}
