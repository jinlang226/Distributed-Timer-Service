package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

type Acceptor struct {
	lis net.Listener
	// 服务器 id
	id int
	// 接受者承诺的提案编号，如果为 0 表示接受者没有收到过任何 Prepare 消息
	promiseNumber int
	// 接受者已接受的提案编号，如果为 0 表示没有接受任何提案
	acceptedNumber int
	// 接受者已接受的提案的值，如果没有接受任何提案则为 nil
	acceptedValue *WriteDataByLine
	// 学习者 id 列表
	learners []int
}

func newAcceptor(id int, learners []int) *Acceptor {
	acceptor := &Acceptor{
		id:       id,
		learners: learners,
	}
	acceptor.server()
	return acceptor
}

func (a *Acceptor) Prepare(args *PaxosMsgArgs, reply *PaxosMsgReply) error {
	if args.Number > a.promiseNumber {
		a.promiseNumber = args.Number
		reply.Number = a.acceptedNumber
		reply.Value = a.acceptedValue
		reply.Ok = true
	} else {
		reply.Ok = false
	}
	return nil
}

func (a *Acceptor) Accept(args *PaxosMsgArgs, reply *PaxosMsgReply) error {
	if args.Number >= a.promiseNumber {
		a.promiseNumber = args.Number
		a.acceptedNumber = args.Number
		a.acceptedValue = args.Value
		reply.Ok = true

		fmt.Println("accept value: ", args.Value)
		_, existed := TW.finishedTasks.Load(args.Value.TaskId)
		if !existed {
			fmt.Println("write to csv during paxos!!!!!!!")
			writeCsvByLine(Filepath+logFilename, args.Value)
			WriteToMap(args.Value.TaskId)
		}
		fmt.Println("already existed, ", time.Now().Format(Format))
		// 后台转发接受的提案给学习者
		//for _, lid := range a.learners {
		//	go func(learner int) {
		//		//todo change the address
		//		addr := fmt.Sprintf("127.0.0.1:%d", learner)
		//		args.From = a.id
		//		args.To = learner
		//		resp := new(PaxosMsgReply)
		//		ok := call(addr, "Learner.Learn", args, resp)
		//		if !ok {
		//			return
		//		}
		//	}(lid)
		//}
	} else {
		reply.Ok = false
	}
	return nil
}

func (a *Acceptor) server() {
	rpcs := rpc.NewServer()
	rpcs.Register(a)
	addr := fmt.Sprintf(":6%d", a.id)
	l, e := net.Listen("tcp", addr)
	//l, e := net.Listen("tcp", fmt.Sprintf(":%s", port2))
	if e != nil {
		log.Fatal("listen error 3:", e)
	}
	a.lis = l
	go func() {
		for {
			conn, err := a.lis.Accept()
			if err != nil {
				continue
			}
			go rpcs.ServeConn(conn)
		}
	}()
}

// 关闭连接
func (a *Acceptor) close() {
	a.lis.Close()
}
