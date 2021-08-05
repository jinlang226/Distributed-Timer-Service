package main

// 启动接受者和学习者 RPC 服务
func start(acceptorIds []int, learnerIds []int) ([]*Acceptor, []*Learner) {
	acceptors := make([]*Acceptor, 0)
	for _, aid := range listenIds {
		a := newAcceptor(aid, learnerIds)
		acceptors = append(acceptors, a)
	}

	learners := make([]*Learner, 0)
	//for _, lid := range learnerIds {
	//	l := newLearner(lid, acceptorIds)
	//	learners = append(learners, l)
	//}

	return acceptors, learners
}

//// 清除数据
//func clean(p []*Proposer, a []*Acceptor, l []*Learner) {
//	for _, i := range p {
//		i.clean()
//	}
//	for _, i := range a {
//		i.clean()
//	}
//	for _, i := range l {
//		i.clean()
//	}
//}

func cleanup(acceptors []*Acceptor, learners []*Learner) {
	for _, a := range acceptors {
		a.close()
	}

	for _, l := range learners {
		l.close()
	}
}

func StartPaxos() ([]*Acceptor, []*Learner, *Proposer) {
	acceptors, learns := start(AcceptorIds, LearnerIds)
	//defer cleanup(acceptors, learns)

	// 1 是提议者 id
	p := &Proposer{
		id:        proposerID,
		acceptors: AcceptorIds,
	}

	return acceptors, learns, p
}
