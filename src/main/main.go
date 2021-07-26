package main

import (
	"fmt"
	"time"
)

var a, l, p = StartPaxos()

func main() {
	//time.Sleep(time.Duration(20) * time.Second)
	timeWheel := CreateTimeWheel(1*time.Second, 60)
	timeWheel.startTW()
	fmt.Println("initialize rpc")
	timeWheel.serverTW()
	fmt.Println("start Batch register")
	BatchRegister()

	//test(timeWheel)
	defer func() {
		for {
		}
	}()
}

func TaskJob(key interface{}) {
	fmt.Println(fmt.Sprintf("%v This is a task job with key: %v", time.Now().Format(Format), key))
}

func test(timeWheel *TimeWheel) {
	fmt.Println(fmt.Sprintf("%v Add task task-5s", time.Now().Format(time.RFC3339)))
	args := &AddTaskArgs{time.Duration(1) * time.Second, 1, time.Now(), TaskJob}
	reply := AddTaskReply{}
	err := timeWheel.AddTask(args, &reply)
	fmt.Println("finish1")
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%v Add task task-10s", time.Now().Format(time.RFC3339)))
	args = &AddTaskArgs{time.Duration(10) * time.Second, 2, time.Now(), TaskJob}
	reply = AddTaskReply{}
	err = timeWheel.AddTask(args, &reply)
	fmt.Println("finish2")
	if err != nil {
		panic(err)
	}
}
