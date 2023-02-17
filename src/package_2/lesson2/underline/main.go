package main

import "fmt"

func foo() (int, string) {
	return 10, "111"
}

func main() {
	//buf := make([]byte, 1024)
	//f, _ := os.Open("/Users/michaelchan/Downloads/【优e出行-36.91元-1个行程】高德打车电子行程单.PDF")
	//defer f.Close()
	//for {
	//	n, _ := f.Read(buf)
	//	if n == 0 {
	//		break
	//
	//	}
	//	os.Stdout.Write(buf[:n])
	//}

	list := []string{"tom", "ally"}

	for _, name := range list {
		fmt.Println(name)
	}

	var (
		a int
		b string
		c float64
	)
	fmt.Println(a, b, c)

	x, _ := foo()
	fmt.Println(x)

}
