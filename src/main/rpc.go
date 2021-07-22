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
	stupid int
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

// This method starts a RPC server for tw
func (tw *TimeWheel) serverTW() {
	rpcs := rpc.NewServer()
	rpcs.Register(tw)
	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Error("listen error 1:", err)
	}
	tw.lis = lis
	go func() {
		for {
			conn, err := tw.lis.Accept()
			if err != nil {
				continue
			}
			go rpcs.ServeConn(conn)
		}
	}()
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