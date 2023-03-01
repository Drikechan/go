package main

import (
	"fmt"
	"time"
)

func main() {
	go hello()
	fmt.Println("main")
	time.Sleep(time.Second)
}

func hello() {
	fmt.Println("hello")
}

// 打印结果是 main hello
