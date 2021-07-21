package main

import (
	"fmt"
	"time"
)

var TW *TimeWheel
var p Proposer

func main() {

	TW = startServer()
	_, _, p := StartPaxos()
	p.id = 1
	InitializeRPC()
	BatchRegister()
}

func startServer() *TimeWheel {
	//初始化一个时间间隔是1s，一共有60个齿轮的时间轮盘，默认轮盘转动一圈的时间是60s
	tw := GetTimeWheel(1*time.Second, 60)
	tw.Start() // start tw
	fmt.Println("start doing tasks")
	return tw
}

func TaskJob(key interface{}) {
	fmt.Println(fmt.Sprintf("%v This is a task job with key: %v", time.Now().Format(Format), key))
}
