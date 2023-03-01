package main

import "fmt"

func testFunction(fn func() int) int {
	return fn()
}

type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

func main() {
	//a := 10
	//
	//b := &a
	//
	//fmt.Println(b, a)
	//
	//fmt.Println(test(1, 2, "3"))
	s1 := testFunction(func() int {
		return 100
	})

	fmt.Println(s1)
}

func test(x, y int, s string) (int, string) {
	n := x + y
	return n, fmt.Sprintf(s, n)
}
