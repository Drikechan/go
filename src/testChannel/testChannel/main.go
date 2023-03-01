package main

import "fmt"

func main() {
	ch4 := make(chan int, 1)
	ch4 <- 10
	fmt.Println(<-ch4)

	close(ch4)

}
