package main

import (
	"fmt"
	"time"
)

var p Proposer

func main() {
	timeWheel := CreateTimeWheel(1*time.Second, 60)
	timeWheel.startTW()
	//_, _, p := StartPaxos()
	p.id = 1
	fmt.Println("initialize rpc")
	//timeWheel.serverTW()
	fmt.Println("start Batch register")
	//BatchRegister()

	fmt.Println(fmt.Sprintf("%v Add task task-5s", time.Now().Format(time.RFC3339)))
	args := &AddTaskArgs{time.Duration(5) * time.Second, 100, time.Now(), TaskJob}
	reply := AddTaskReply{}
	err := timeWheel.AddTask(args, &reply)
	fmt.Println("finish")
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%v Add task task-10s", time.Now().Format(time.RFC3339)))
	args = &AddTaskArgs{time.Duration(10) * time.Second, 200, time.Now(), TaskJob}
	reply = AddTaskReply{}
	err = timeWheel.AddTask(args, &reply)
	fmt.Println("finish")
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%v Add task task-20s", time.Now().Format(time.RFC3339)))
	args = &AddTaskArgs{time.Duration(20) * time.Second, 300, time.Now(), TaskJob}
	reply = AddTaskReply{}
	err = timeWheel.AddTask(args, &reply)
	fmt.Println("finish")
	if err != nil {
		panic(err)
	}
}

func TaskJob(key interface{}) {
	fmt.Println(fmt.Sprintf("%v This is a task job with key: %v", time.Now().Format(Format), key))
}
