package main

import (
	"fmt"
	"time"
)

// receive tasks from Client, parameter from client might be string, need parsing interval and uuid from it.
func Register(interval time.Duration, uuid int) {
	register(interval, uuid)
	//rpc to other servers
	args := RPCBackupArgs{
		Interval: interval,
		Uuid:     uuid,
	}
	backup(args)
}

// scan the csv
func BatchRegister() {
	result, err := ReadFile("/root/ft_local/Distributed-Timer-Service/src/test/test.csv")
	if err != nil {
		fmt.Println("err in read file")
	}
	//for each line in csv data structure:
	for _, tasks := range result {
		for _, items := range tasks {
			duration := time.Duration(items[1]) * time.Second
			uuid := int(items[0])
			Register(duration, uuid)
		}
	}
}

// rpc calling, public method
// finish tasks from other server
func Backup(args *RPCBackupArgs, reply *RPCBackupReply) error {
	uuid := args.Uuid
	interval := args.Interval
	register(interval, uuid)
	reply.Msg = "Backup Succeed.\n"
	reply.Code = 0
	return nil
}

// register tasks
func register(interval time.Duration, uuid int) {
	fmt.Println(fmt.Sprintf("%s Add Task ID: %d", time.Now().Format(Format), uuid))
	err := TW.AddTask(interval, uuid, time.Now(), TaskJob)
	if err != nil {
		panic(err)
	}
}

func backup(args RPCBackupArgs) {
	reply := RPCBackupReply{}
	// call method is defined in
	for _, socketName:= range SocketNames {
		if socketName != LocalName {
			if ok := callRPC(socketName, "Backup", args, &reply); !ok {
				fmt.Printf("Register: backup register error\n")
			}
		}
	}

}
