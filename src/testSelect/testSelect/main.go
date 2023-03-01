package main

import (
	"fmt"
	"time"
)

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "4"
}

func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "5"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go test1(ch1)
	go test2(ch2)

	select {
	case s1 := <-ch1:
		fmt.Println(s1, "s1")
	case s2 := <-ch2:
		fmt.Println(s2, "s2")
	}

}
