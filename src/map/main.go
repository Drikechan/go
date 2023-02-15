package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var p1 map[int]string
	p1 = make(map[int]string)
	p1[1] = "tom"
	p1[2] = "michael"
	fmt.Println(p1)

	var p2 map[int]string = map[int]string{}
	p2[1] = "tom"
	fmt.Println(p2)

	var p3 map[int]string = make(map[int]string)
	p3[1] = "kobe"
	p3[2] = "brant"
	fmt.Println(p3)

	p4 := map[int]string{}
	p4[1] = "kiwi"
	fmt.Println(p4)

	p5 := make(map[int]string)
	p5[1] = "cai"
	fmt.Println(p5)

	p6 := map[int]string{
		1: "Nets",
	}

	fmt.Println(p6)

	generateJson()
	editDelete()
}

func editDelete() {
	person := map[int]string{
		1: "Tom",
		2: "Aaron",
		3: "John",
	}

	fmt.Println(person)

	delete(person, 2)

	fmt.Println(person)

	person[2] = "kill"
	fmt.Println(person)
}

func generateJson() {
	res := make(map[string]interface{})
	res["code"] = 200
	res["data"] = map[string]interface{}{
		"username": "Jack",
		"age":      "10",
	}

	// 序列化
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println(errs)
	}
	fmt.Println(string(jsons))

	// 反序列化
	res2 := make(map[string]interface{})
	errs = json.Unmarshal([]byte(jsons), &res2)

	if errs != nil {
		fmt.Println(errs)
	}

	fmt.Println(res2)
}
