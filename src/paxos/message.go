package paxos

import (
	"modu/src/message"
	"net/rpc"
)

type MsgArgs struct {
	Number int                      // 提案编号
	Value  *message.WriteDataByLine // 提案的值
	From   int                      // 发送者 id
	To     int                      // 接收者 id
}

type MsgReply struct {
	Ok     bool
	Number int
	Value  interface{}
}

func call(srv string, name string, args interface{}, reply interface{}) bool {
	c, err := rpc.Dial("tcp", srv)
	if err != nil {
		return false
	}
	defer c.Close()

	err = c.Call(name, args, reply)
	if err == nil {
		return true
	}
	return false
}
