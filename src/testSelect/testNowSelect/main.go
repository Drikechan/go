package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "6"
	}()

	go func() {
		ch2 <- "7"
	}()

	select {
	case s1 := <-ch1:
		fmt.Println("s1", s1)
	case s2 := <-ch2:
		fmt.Println("s2", s2)
	}

}
