package main

import (
	"log"
	"net"
	"net/rpc"
)

type Learner struct {
	lis         net.Listener
	id          int                  // 学习者 id
	acceptedMsg map[int]PaxosMsgArgs // 记录接受者已接受的提案：[接受者 id]请求消息
}

func newLearner(id int, acceptorIds []int) *Learner {
	learner := &Learner{
		id:          id,
		acceptedMsg: make(map[int]PaxosMsgArgs),
	}
	for _, aid := range acceptorIds {
		learner.acceptedMsg[aid] = PaxosMsgArgs{
			Number: 0,
			Value:  nil,
		}
	}
	learner.server(id)
	return learner
}

func (l *Learner) Learn(args *PaxosMsgArgs, reply *PaxosMsgReply) error {
	a := l.acceptedMsg[args.From]
	if a.Number < args.Number {
		l.acceptedMsg[args.From] = *args
		reply.Ok = true

	} else {
		reply.Ok = false
	}
	return nil
}

//func (l *Learner) Chosen() interface{} {
//	acceptCounts := make(map[int]int)
//	acceptMsg := make(map[int]PaxosMsgArgs)
//
//	for _, accepted := range l.acceptedMsg {
//		if accepted.Number != 0 {
//			acceptCounts[accepted.Number]++
//			acceptMsg[accepted.Number] = accepted
//		}
//	}
//
//	for n, count := range acceptCounts {
//		if count >= l.majority() {
//			//todo save value in logs and local map
//			return acceptMsg[n].Value
//		}
//	}
//	return nil
//}

func (l *Learner) majority() int {
	return len(l.acceptedMsg)/2 + 1
}

func (l *Learner) server(id int) {
	rpcs := rpc.NewServer()
	rpcs.Register(l)
	lis, e := net.Listen("tcp", ":8007")
	if e != nil {
		log.Fatal("listen error 4:", e)
	}
	l.lis = lis
	go func() {
		for {
			conn, err := l.lis.Accept()
			if err != nil {
				continue
			}
			go rpcs.ServeConn(conn)
		}
	}()
}

// 关闭连接
func (l *Learner) close() {
	l.lis.Close()
}
