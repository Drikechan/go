package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type Person struct {
	Name string
	Age  int
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	var p1 Person
	p1.Name = "tom"
	p1.Age = 30
	fmt.Println(p1)

	var p2 = Person{Name: "michael", Age: 10}
	fmt.Println(p2)

	p3 := Person{Name: "jordan", Age: 60}
	fmt.Println(p3)

	p4 := struct {
		Name string
		Age  int
	}{Name: "kobe", Age: 43}
	fmt.Println(p4)

	// 生成json
	var res Result
	res.Code = 200
	res.Message = "成功"
	// 序列化
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println(errs)
	}
	fmt.Println(string(jsons))

	// 反序列化

	var res2 Result
	errs = json.Unmarshal(jsons, &res2)
	fmt.Println(res2)

	go func() {
		var a sync.WaitGroup
		a.Add(1)
		getTerminalCode()
	}()
}

func getTerminalCode() {
	time.Sleep(time.Second * 5)
	newArrayList := []int{}
	for i := 0; i < 10; i++ {
		newArrayList[i] = i
		fmt.Println(i)
	}

}
