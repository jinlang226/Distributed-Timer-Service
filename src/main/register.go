package main

import (
	"fmt"
	"modu/src/timeWheel"
	"time"
	"modu/src/network"
)

// receive tasks from Client, parameter from client might be string, need parsing interval and uuid from it.
func Register(interval time.Duration, uuid int) {
	register(interval, uuid)
	//rpc to other servers
	args := network.RPCBackupArgs {
		Interval: interval,
		Uuid: uuid,
	}
	backup(args)
}

// scan the csv
func BatchRegister(tw *timeWheel.TimeWheel) {
	//for each line in csv data structure:
	//for line in csv  {
	//uuid, interval need to be split from line
	interval := 2 * time.Second
	uuid := 0
	Register(interval, uuid)
	//}
}

// rpc calling, public method
// finish tasks from other server
func Backup(args *network.RPCBackupArgs, reply *network.RPCBackupReply) error {
	uuid := args.Uuid
	interval := args.Interval
	register(interval, uuid)
	reply.Msg = "Backup Succeed.\n"
	reply.Code = 0
	return nil
}

// regist tasks
func register(interval time.Duration, uuid int) {
	fmt.Println(fmt.Sprintf("%s Add Task ID: %d", time.Now().Format(format), uuid))
	err := tw.AddTask(interval, uuid, time.Now(), TaskJob)
	if err != nil {
		panic(err)
	}
}

func backup(args network.RPCBackupArgs) {
	reply := network.RPCBackupReply{}
	// call method is defined in
	if ok := network.Call(network.Socketname1, "Backup", args, &reply); !ok {
		fmt.Printf("Register: backup register error\n")
	}
}
