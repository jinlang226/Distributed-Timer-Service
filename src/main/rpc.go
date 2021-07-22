package main

import (
	"fmt"
	"github.com/go-playground/log"
	"net"
	"net/rpc"
	"time"
)

type AddTaskArgs struct {
	interval time.Duration
	uuid     int
	execTime time.Time
	taskJob  interface{}
}

type AddTaskReply struct {
}

type BackupArgs struct {
	Interval time.Duration
	Uuid     int
}

type BackupReply struct {
	Msg  string
	Code int
}

type PaxosMsgArgs struct {
	Number int              // 提案编号
	Value  *WriteDataByLine // 提案的值
	From   int              // 发送者 id
	To     int              // 接收者 id
}

type PaxosMsgReply struct {
	Ok     bool
	Number int
	Value  *WriteDataByLine
}

// This method starts a RPC server
func InitializeRPC() {
	rpc.Register(TW)
	l, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Error("listen error 1:", err)
	}
	conn, err := l.Accept()
	if err != nil {
		log.Error("listen error 2:", err)
	}
	rpc.ServeConn(conn)
}

//
// send an RPC request to other servers, wait for the response.
// usually returns true.
// returns false if something goes wrong.
//
func call(sockname string, rpcname string, args interface{}, reply interface{}) bool {
	c, err := rpc.Dial("tcp", sockname+":80")
	if err != nil {
		log.Error("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	fmt.Printf("calling:%s::%s() error: %s\n", sockname, rpcname, err)
	return false
}