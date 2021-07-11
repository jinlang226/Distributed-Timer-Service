package main

import (
	"fmt"
	"modu/src/timeWheel"
	"time"
	"encoding/csv"
	"io"
	"log"
	"strings"
)

type RPCBackupArgs struct {
	Interval time.Duration
	Uuid int
}

type RPCBackupReply struct {
	Msg string
	Code int
}

// receive tasks from Client, parameter from client might be string, need parsing interval and uuid from it.
func Register() {
	
	regist(tw, interval, uuid)
	//rpc to other servers
	backup({interval, uuid})

}

// scan the csv
func BtachRegister(tw *timeWheel.TimeWheel) {
	//for each line in csv data structure:
	//for line in csv  {
	//uuid, interval need to be split from line
		interval := 2 * time.Second
		uuid := 0
		Register()


	//}
}

// Backup with capital letter 'B' is used for rpc calling.
func Backup(args *RPCBackupArgs, reply *RPCBackupReply) error {
	uuid := args.Uuid;
	interval := args.Interval
	register(tw *timeWheel.TimeWheel, interval, uuid int)
	reply.Msg = "Backup Succeed.\n"
	reply.Code = 0
	return nil
}

// regist tasks
func register(tw *timeWheel.TimeWheel, interval time.Duration, uuid int) {
	fmt.Println(fmt.Sprintf("%s Add Task ID: %d", time.Now().Format(format), uuid))
	err := tw.AddTask(interval, uuid, time.Now(), TaskJob)
	if err != nil {
		panic(err)
	}
}

func backup(args RPCParameters) {
	reply := RPCBackupReply{}
	// call method is defined in
	if ok := call(sockname, "Scheduler.Backup", args, &reply); !ok {
		fmt.Printf("Register: backup register error\n")
	}
}

