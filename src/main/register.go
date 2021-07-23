package main

import (
	"fmt"
	"github.com/go-playground/log"
	"strconv"
	"time"
)

// receive tasks from Client, parameter from client might be string, need parsing interval and uuid from it.
func Register(interval time.Duration, uuid int) {
	register(interval, uuid)
	//rpc to other servers
	args := BackupArgs{
		Interval: interval,
		Uuid:     uuid,
	}
	backup(args)
}

// scan the csv
func BatchRegister() {
	//result, err := ReadFile("/root/ft_local/Distributed-Timer-Service/src/test/test.csv")
	//only for testing
	result, err := ReadFile(Filepath+FilenameRead)
	if err != nil {
		fmt.Println("err in read file")
	}
	//for each line in csv data structure:
	for _, items := range result {
		//fmt.Println(items)
		d, err := strconv.Atoi(items[1])
		if err != nil {
			log.Error(err)
		}
		duration := time.Duration(d) * time.Second
		uuid, err := strconv.Atoi(items[0])
		if err != nil {
			log.Error(err)
		}
		fmt.Println("duration: ", duration, " uuid: ", uuid)
		Register(duration, uuid)
		//}
	}
}

// rpc calling, public method
// finish tasks from other server
func Backup(args *BackupArgs, reply *BackupReply) error {
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
	args := &AddTaskArgs{interval, uuid, time.Now(), TaskJob}
	reply := AddTaskReply{}
	err := TW.AddTask(args, &reply)
	//fmt.Println(reply.stupid)
	if err != nil {
		panic(err)
	}
}

func backup(args BackupArgs) {
	reply := BackupReply{}
	// call method is defined in
	//for _, socketName := range SocketNames {
	//	if socketName != LocalName {
	//		if ok := call(socketName, "Backup", args, &reply); !ok {
	//			fmt.Printf("Register: backup register error\n")
	//		}
	//	}
	//}
	if ok := call(Socketname2, "Backup", args, &reply); !ok {
		fmt.Printf("Register: backup register error\n")
	}
	if ok := call(Socketname3, "Backup", args, &reply); !ok {
		fmt.Printf("Register: backup register error\n")
	}
}
