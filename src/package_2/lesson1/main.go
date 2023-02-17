package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.4, 32.5
	fmt.Println(fToc(freezingF))
	fmt.Println(fToc(boilingF))
	fmt.Println(calculationNumber())
	fmt.Println(toAdd())
}

func fToc(f float64) float64 {
	return (f - 32) * 5 / 9
}

func calculationNumber() float64 {
	return 0.14 * 100
}

func toAdd() float64 {
	return 0.1 + 0.2
}
