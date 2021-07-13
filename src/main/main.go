package main

import (
	"fmt"
	"github.com/go-playground/log"
	"modu/src/paxos"
	"modu/src/timeWheel"
	"time"
)

func main() {
	tw := startServer()
	_, learns, p := paxos.StartPaxos()

	value := p.Propose("hello world")
	if value != "hello world" {
		log.Error("value = %s, excepted %s", value, "hello world")
	}

	learnValue := learns[0].Chosen()
	if learnValue != value {
		log.Error("learnValue = %s, excepted %s", learnValue, "hello world")
	}

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
