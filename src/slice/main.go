package main

import "fmt"

func main() {
	// 声明切片
	var sliceOne []int
	sliceOne = append(sliceOne, 7)
	fmt.Println(len(sliceOne))

	sliceTwo := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sliceTwo[:len(sliceTwo)-2])

	fmt.Println(sliceTwo[2:])

	fmt.Println(sliceTwo)
}
