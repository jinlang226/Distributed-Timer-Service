package main

import (
	"fmt"
	"modu/src/timeWheel"
	"time"
)

// regist tasks
func regist(tw *timeWheel.TimeWheel, interval time.Duration, uuid int) {
	fmt.Println(fmt.Sprintf("%s Add Task ID: %d", time.Now().Format(format), uuid))
	err := tw.AddTask(interval, uuid, time.Now(), TaskJob)
	if err != nil {
		panic(err)
	}
}

// receive tasks from other server
func Regist() {
	//rpc to other servers
	

}

func RPCregist() {
	//
}

// scan the csv
func BtachRegist(tw *timeWheel.TimeWheel) {
	//for each line in csv data structure:
	//for line in csv  {
	//uuid, interval need to be split from line
		interval := 2 * time.Second
		uuid := 0
		regist(tw, interval, uuid)


	//}
}
