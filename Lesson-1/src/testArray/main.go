package main

import "fmt"

func main() {
	myArr := [5]int{1, 2, 3, 4, 5}
	mySlice := myArr[1:3]
	fmt.Printf("mySlice %+v\n", mySlice)
	fullSlice := myArr[:]
	removeThreeItem := deleteItem(fullSlice, 2)
	fmt.Printf("mySlice %+v\n", removeThreeItem)
	print(1)
}
