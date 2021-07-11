package main

import (
	"fmt"
	"modu/src/timeWheel"
	"time"
)
func main() {
	startTW()
}

func TaskJob(key interface{}) {
	fmt.Println(fmt.Sprintf("%v This is a task job with key: %v", time.Now().Format(time.RFC3339), key))
	//put it into log
}

func startTW() {
	//初始化一个时间间隔是1s，一共有60个齿轮的时间轮盘，默认轮盘转动一圈的时间是60s
	tw := timeWheel.GetTimeWheel(1*time.Second, 60)

	// 启动时间轮盘
	tw.Start()

	if tw.IsRunning() {
		//for each line in csv data structure:
		//for uuid, interval in each line:
		uuid := 0
		interval := 2 * time.Second
		fmt.Println(fmt.Sprintf("%s Add Task ID: %d", time.Now().Format("2006/1/2 15:04:05"), uuid))
		err := tw.AddTask(interval, uuid, time.Now(), TaskJob)
		//rpc calls to at least two other servers
		if err != nil {
			panic(err)
		}
	} else {
		panic("TimeWheel is not running")
	}

	// 删除task
	//fmt.Println("Remove task task-5s")
	//err := tw.RemoveTask("task-5s")
	//if err != nil {
	//	panic(err)
	//}


	fmt.Println("finished tasks")
}
