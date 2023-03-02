package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	// \d是数字
	reQQEmail = `(\d+)@qq.com`
)

func GetEmail() {
	res, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	HandleError(err, "Http get err")
	defer res.Body.Close()

	pageBytes, err := ioutil.ReadAll(res.Body)

	HandleError(err, "ReadAll error")
	convertStr := string(pageBytes)

	reg := regexp.MustCompile(reQQEmail)

	result := reg.FindStringSubmatch(convertStr)

	for _, value := range result {
		fmt.Println("email", value)
		fmt.Println("qq", value)
	}
}

func HandleError(err error, errorResult string) {
	if err != nil {
		fmt.Println(err, errorResult)
	}
}

func main() {
	GetEmail()
}
