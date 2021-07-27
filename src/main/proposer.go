package main

import (
	"fmt"
	"github.com/go-playground/log"
)

type Proposer struct {
	// 服务器 id
	id int
	// 当前提议者已知的最大轮次
	round int
	// 提案编号 = (轮次, 服务器 id)
	number int
	// 接受者 id 列表
	acceptors []int
}

func (p *Proposer) Propose(v *WriteDataByLine) interface{} {
	fmt.Println("paxos propose")
	p.round++
	p.number = p.proposalNumber()

	// 第一阶段(phase 1)
	prepareCount := 0
	maxNumber := 0
	for _, aid := range p.acceptors {
		args := PaxosMsgArgs{
			Number: p.number,
			From:   p.id,
			To:     aid,
			Value:  v,
		}
		reply := PaxosMsgReply{}

		err := call(SocketNames[aid], "Prepare", args, &reply)
		if !err {
			continue
		} else {
			log.Error("Acceptor.Prepare error: %v", err)
		}

		if reply.Ok {
			prepareCount++
			if reply.Number > maxNumber {
				maxNumber = reply.Number
				v = reply.Value
			}
		}

		if prepareCount == p.majority() {
			break
		}
	}

	// 第二阶段(phase 2)
	acceptCount := 0
	if prepareCount >= p.majority() {
		for _, aid := range p.acceptors {
			args := PaxosMsgArgs{
				Number: p.number,
				Value:  v,
				From:   p.id,
				To:     aid,
			}
			reply := PaxosMsgReply{}
			//todo change the address
			err := call(SocketNames[aid], "Accept", args, &reply)
			if !err {
				continue
			} else {
				log.Error("Acceptor.Prepare error: %v", err)
			}

			if reply.Ok {
				acceptCount++
			}
		}
	}

	if acceptCount >= p.majority() {
		// 选择的提案的值
		//todo save locally
		return v
	}
	return nil
}

func (p *Proposer) majority() int {
	return len(p.acceptors)/2 + 1
}

// 提案编号 = (轮次, 服务器 id)
func (p *Proposer) proposalNumber() int {
	return p.round<<16 | p.id
}
