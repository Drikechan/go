package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, _ := http.Get("127.0.0.1:8000")
	fmt.Println(res)

	defer res.Body.Close()

	buf := make([]byte, 1024)

	for {
		n, err := res.Body.Read(buf)

		if err != nil && err != io.EOF {
			return
		} else {
			res := string(buf[:n])
			fmt.Println(res)
			return
		}
	}
}
