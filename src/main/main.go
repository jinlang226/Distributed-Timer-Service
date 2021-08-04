package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-playground/log/v7"
	"github.com/go-playground/log/v7/handlers/console"
)

var a, l, p = StartPaxos()

func main() {
	//initialize logs
	cLog := console.New(true)
	log.AddHandler(cLog, log.AllLevels...)

	// Trace
	defer log.WithTrace().Info("time to run")

	//delete old log file
	err := os.Remove(Filepath + logFilename) //删除文件test.txt
	if err != nil {
		//如果删除失败则输出 file remove Error!
		fmt.Println("file remove Error!")
		//输出错误详细信息
		fmt.Printf("%s", err)
	} else {
		//如果删除成功则输出 file remove OK!
		fmt.Print("file remove OK!")
	}

	//time.Sleep(time.Duration(20) * time.Second)
	timeWheel := CreateTimeWheel(1*time.Second, 60)
	timeWheel.startTW()
	log.Info("initialize rpc")
	timeWheel.serverTW()
	log.Info("start Batch register")
	time.Sleep(time.Duration(5) * time.Second)
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
