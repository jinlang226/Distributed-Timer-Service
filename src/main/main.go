package main

import (
	"fmt"
	TimeWheel "mod/src/timeWheel"
	"time"
)

func TaskJob(key interface{}) {
	fmt.Println(fmt.Sprintf("%v This is a task job with key: %v", time.Now().Format(time.RFC3339), key))
	//put it into log
}

func startTW() {
	//初始化一个时间间隔是1s，一共有60个齿轮的时间轮盘，默认轮盘转动一圈的时间是60s
	tw := TimeWheel.GetTimeWheel(1*time.Second, 60)

	// 启动时间轮盘
	tw.Start()

	if tw.IsRunning() {
		//for each line in csv data structure:
		//for uuid, interval in :
		uuid := 0
		interval := 2 * time.Second
		fmt.Println(fmt.Sprintf("%v Add task id: %d", time.Now().Format(time.RFC3339)), uuid)
		err := tw.AddTask(interval, uuid, time.Now(), -1, TaskJob)
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

	// 关闭时间轮盘
	if tw.Finished() == true {
		tw.Stop()
	}
	fmt.Println("finished tasks")
}
