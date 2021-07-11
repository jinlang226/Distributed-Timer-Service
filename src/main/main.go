package main

import (
	"fmt"
	"modu/src/timeWheel"
	"time"
)

const format string = "2006/1/2 15:04:05"

func main() {
	//初始化一个时间间隔是1s，一共有60个齿轮的时间轮盘，默认轮盘转动一圈的时间是60s




}

func TaskJob(key interface{}) {
	fmt.Println(fmt.Sprintf("%v This is a task job with key: %v", time.Now().Format(format), key))
}

func startTW(tw *timeWheel.TimeWheel) {
	tw.Start()	// start tw
	fmt.Println("start doing tasks")
}

func Register(tw *timeWheel.TimeWheel) {
	if tw.IsRunning() {
		//1. for each line in csv data structure:
		//for line in csv  {
			//uuid, interval need to be split from line
			uuid := 0
			interval := 2 * time.Second
			fmt.Println(fmt.Sprintf("%s Add Task ID: %d", time.Now().Format(format), uuid))
			err := tw.AddTask(interval, uuid, time.Now(), TaskJob)
			//rpc calls to at least two other servers
		//}

		//2. receive tasks from peers




		if err != nil {
			panic(err)
		}
	} else {
		panic("TimeWheel is not running")
	}

}

func startServer() {
	tw := timeWheel.GetTimeWheel(1*time.Second, 60)
	startTW(tw)



}