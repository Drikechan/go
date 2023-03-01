package main

import (
	"fmt"
	"time"
)

func main() {
	//timer1 := time.NewTimer(time.Second * 2)
	//fmt.Println(timer1)
	//t1 := time.Now()
	//fmt.Println(t1)
	//
	//t2 := <-timer1.C
	//fmt.Println(t2)

	//timer2 := time.NewTimer(time.Second)
	//for  {
	//	<-timer2.C
	//	fmt.Println("时间到")
	//}

	//time.Sleep(time.Second)
	//
	//time3 := time.NewTimer(2 * time.Second)
	//<-time3.C
	//fmt.Println("2s")
	//<-time.After(2 * time.Second)
	//fmt.Println("again 2s")

	// 停止定时器
	//time4 := time.NewTimer(2 * time.Second)
	//go func() {
	//	<-time4.C
	//	fmt.Println("start")
	//}()
	//b := time4.Stop()
	//if b {
	//	fmt.Println("关闭了")
	//}

	//time5 := time.NewTimer(3 * time.Second)
	//time5.Reset(1 * time.Second)
	//fmt.Println(time.Now())
	//fmt.Println(<-time5.C)
	//for  {
	//
	//}

	//time5 := time.Now()
	//fmt.Println(time5.Year())
	//fmt.Println(time5.Month().String())
	//fmt.Println(time5.Unix())
	//fmt.Println(time5.Hour())
	//fmt.Println(time5.Add(time.Hour))
	//fmt.Println(time5.Format("2006-01-02 03:04:05.000 PM Mon Jan"))

	// 定时器
	//time6 := time.Tick(4 * time.Second)
	//for i := range time6{
	//	fmt.Println(i)
	//}

	time7 := time.NewTicker(1 * time.Second)
	i := 0
	go func() {
		for {
			i++
			fmt.Println(<-time7.C)
			if i == 5 {
				time7.Stop()
			}
		}
	}()
	for {

	}
}
