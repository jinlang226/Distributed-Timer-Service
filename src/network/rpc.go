package network

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
	"time"
)

const (
	Socketname1 string = "IP1"
	Socketname2 string = "IP2"
	Socketname3 string = "IP3"
	Socketname4 string = "IP4"
	Socketname5 string = "IP5"
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
func masterSock() string {
	s := "/var/tmp/824-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}

// This method starts a RPC server
func (m *Master) InitializeRPC() {
	rpc.Register(m)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := masterSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

//
// send an RPC request to the master, wait for the response.
// usually returns true.
// returns false if something goes wrong.
//
func Call(sockname string, rpcname string, args interface{}, reply interface{}) bool {
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	fmt.Printf("calling:%s::%s() error: %s\n", sockname, rpcname, err)
	return false
}