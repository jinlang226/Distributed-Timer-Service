package paxos_ref

import (
	"fmt"
	"github.com/go-playground/log/v7"
)

// 启动接受者和学习者 RPC 服务
func start(acceptorIds []int, learnerIds []int) ([]*Acceptor, []*Learner) {
	acceptors := make([]*Acceptor, 0)
	for _, aid := range acceptorIds {
		a := newAcceptor(aid, learnerIds)
		acceptors = append(acceptors, a)
	}

	learners := make([]*Learner, 0)
	for _, lid := range learnerIds {
		l := newLearner(lid, acceptorIds)
		learners = append(learners, l)
	}

	return acceptors, learners
}

func cleanup(acceptors []*Acceptor, learners []*Learner) {
	for _, a := range acceptors {
		a.close()
	}

	for _, l := range learners {
		l.close()
	}
}

func (p *Proposer) clean() {
	p.round = 0
	p.id = 0
	p.number = 0
}

func (a *Acceptor) clean() {
	a.promiseNumber = 0
	a.acceptedValue = 0
	a.acceptedValue = nil
}

func (le *Learner) clean() {
	le.acceptedMsg = nil
}

// 清除数据
func clean(p Proposer, a []*Acceptor) {
	p.clean()
	for _, i := range a {
		i.clean()
	}
}

func TestSingleProposer() {
	// 1001, 1002, 1003 是接受者 id
	acceptorIds := []int{1001, 1002, 1003}
	// 2001 是学习者 id
	learnerIds := []int{2001}
	acceptors, learns := start(acceptorIds, learnerIds)

	defer cleanup(acceptors, learns)

	// 1 是提议者 id
	p := &Proposer{
		id:        1,
		acceptors: acceptorIds,
	}

	value := p.propose("hello world")
	if value != "hello world" {
		log.Error("value = %s, excepted %s", value, "hello world")
	}

	learnValue := learns[0].chosen()
	if learnValue != value {
		log.Error("learnValue = %s, excepted %s", learnValue, "hello world")
	}
}

func TestTwoProposers() {
	// 1001, 1002, 1003 是接受者 id
	acceptorIds := []int{1001, 1002, 1003}
	//acceptorIds1 :=[]int{1004, 1002, 1003}
	// 2001 是学习者 id
	learnerIds := []int{2001}
	acceptors, learns := start(acceptorIds, learnerIds)
	defer cleanup(acceptors, learns)

	// 1, 2 是提议者 id
	p1 := Proposer{
		id:        1,
		acceptors: acceptorIds,
	}
	v1 := p1.propose("hello world")
	fmt.Println("!!!!!!!!!!!v1: ", v1)

	clean(p1, acceptors)
	p1.round = 0
	p1.id = 1
	p1.number = 0

	for _, a := range acceptors {
		a.promiseNumber = 0
		a.acceptedValue = 0
		a.acceptedValue = nil
		a.acceptedNumber = 0
	}

	fmt.Println("after clean !123849012378490173491207384291073894107 3849712398")
	v4 := p1.propose("wtf")
	fmt.Println("!!!!!!!!!!!!!!!!!1v4: ", v4)


	learnValue := learns[0].chosen()
	if learnValue != v1 {
		log.Error("learnValue = %s, excepted %s", learnValue, v1)
	}
}
