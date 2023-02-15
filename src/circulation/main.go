package main

import "fmt"

func main() {
	//circulationArray()

	//circulationSlice()

	//circulationMap()

	//testBreak()

	//testContinue()
	p := new(int)
	fmt.Println(*p)
	testGoto()
}

// 循环数组
func circulationArray() {
	company := [3]string{"Amazon", "microsoft", "google"}
	fmt.Println(len(company), company)

	for k, v := range company {
		fmt.Println(k, v)
	}

	for i := range company {
		fmt.Println(i, company[i])
	}

	for i := 0; i < len(company); i++ {
		fmt.Println(i, company[i])
	}

	for _, name := range company {
		fmt.Println(name)
	}
}

func circulationSlice() {
	company := []string{"Amazon", "microsoft", "google"}

	for k, v := range company {
		fmt.Println(k, v)
	}

	for i := range company {
		fmt.Println(i, company[i])
	}

	for i := 0; i < len(company); i++ {
		fmt.Println(i, company[i])
	}

	for _, name := range company {
		fmt.Println(name)
	}
}

// 循环map
func circulationMap() {
	company := map[int]string{
		1: "Amazon",
		2: "microsoft",
		3: "google",
	}

	for k, v := range company {
		fmt.Println(k, v)
	}

	for i := range company {
		fmt.Println(i, company[i])
	}

	for i := 1; i <= len(company); i++ {
		fmt.Println(i, company[i])
	}

	for _, name := range company {
		fmt.Println(name)
	}
}

func testBreak() {
	for i := 0; i < 10; i++ {
		if i == 2 {
			fmt.Println("进来了")
			break
		}
		fmt.Println(i)
	}
}

func testContinue() {
	for i := 0; i < 10; i++ {
		if i == 2 {
			continue
			fmt.Println("continue")
		}
		fmt.Println(i)
	}
}

func testGoto() {
	for i := 0; i < 10; i++ {
		if i == 2 {
			goto TOEND
		}
		fmt.Println(i)
	}

TOEND:
	fmt.Println("打乱顺序")
}
