package main

import "fmt"

func main() {
	var arrOne [5]int
	fmt.Println(arrOne)

	var arrTwo = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrTwo)

	arrThree := [5]int{6, 7, 8, 9, 10}
	fmt.Println(arrThree)

	arrFour := [...]int{11, 12, 13, 14, 15}
	fmt.Println(arrFour)
	// 指定数组下标
	arrFive := [5]int{0: 3, 1: 5, 3: 2}
	fmt.Println(arrFive)

	arrSix := [3][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}}
	fmt.Println(arrSix)

	arrSeven := [...][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {11, 12, 13, 14, 15}}
	fmt.Println(arrSeven)
}
