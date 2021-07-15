package main

import (
	"fmt"
	"modu/src/paxos"
	"time"
)

var acceptors []*paxos.Acceptor
var learners []*paxos.Learner
var proposer *paxos.Proposer

func main() {
	startServer()
	acceptors, learners, proposer = paxos.StartPaxos()
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
	fmt.Println(fmt.Sprintf("%v This is a task job with key: %v", time.Now().Format(format), key))
}
