package main

import (
	"fmt"
	"modu/src/timeWheel"
	"time"
)

const format string = "2006/1/2 15:04:05"
var tw *timeWheel.TimeWheel
var filename = "idk"
var filepath = "wholePathName"

func main() {
	tw := startServer()
	BatchRegister(tw)

}

func startServer() *timeWheel.TimeWheel {
	//初始化一个时间间隔是1s，一共有60个齿轮的时间轮盘，默认轮盘转动一圈的时间是60s
	tw := timeWheel.GetTimeWheel(1*time.Second, 60)
	tw.Start() // start tw
	fmt.Println("start doing tasks")
	return tw
}

func TaskJob(key interface{}) {
	fmt.Println(fmt.Sprintf("%v This is a task job with key: %v", time.Now().Format(format), key))
}
