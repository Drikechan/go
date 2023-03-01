package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/go_test", myHandler)

	http.ListenAndServe("127.0.0.1:8000", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	w.Write([]byte("www.baidu.com"))
}
