package main

// 启动接受者和学习者 RPC 服务
func start(acceptorIds []int) ([]*Acceptor) {
	acceptors := make([]*Acceptor, 0)
	for _, aid := range listenIds {
		a := newAcceptor(aid)
		acceptors = append(acceptors, a)
	}

	return acceptors
}

func cleanup(acceptors []*Acceptor) {
	for _, a := range acceptors {
		a.close()
	}
}

func StartPaxos() ([]*Acceptor,*Proposer) {
	acceptors := start(AcceptorIds)
	//defer cleanup(acceptors, learns)

	// 1 是提议者 id
	p := &Proposer{
		id:        proposerID,
		acceptors: AcceptorIds,
	}

	return acceptors, p
}
