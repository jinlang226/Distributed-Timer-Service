package main

import (
	"fmt"
	"github.com/go-playground/log"
	"net"
	"net/rpc"
	"time"
)

type RPCBackupArgs struct {
	Interval time.Duration
	Uuid     int
}

type RPCBackupReply struct {
	Msg  string
	Code int
}

//from MIT 6.824
//func serverSocket() string {
//	s := "/var/tmp/824-mr-"
//	s += strconv.Itoa(os.Getuid())
//	return s
//}

// This method starts a RPC server
func (tw *TimeWheel) InitializeRPC() {
	rpc.Register(tw)
	l, err := net.Listen("tcp", localName)
	if err != nil {
		log.Error("listen error:", err)
	}
	conn, err := l.Accept()
	if err != nil {
		log.Error("listen error:", err)
	}
	rpc.ServeConn(conn)
}

//
// send an RPC request to other servers, wait for the response.
// usually returns true.
// returns false if something goes wrong.
//
func call(sockname string, rpcname string, args interface{}, reply interface{}) bool {
	c, err := rpc.Dial("tcp", sockname)
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
