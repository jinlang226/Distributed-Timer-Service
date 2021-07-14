package paxos

var serverID1 = 1
var serverID2 = 2
var serverID3 = 3
var serverID4 = 4
var serverID5 = 5

var acceptorIds = []int{serverID2, serverID3, serverID4}
var learnerIds = []int{serverID5}

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

func StartPaxos() ([]*Acceptor, []*Learner, *Proposer) {
	acceptors, learns := start(acceptorIds, learnerIds)
	//defer cleanup(acceptors, learns)

	// 1 是提议者 id
	p := &Proposer{
		id:        1,
		acceptors: acceptorIds,
	}

	return acceptors, learns, p
}
