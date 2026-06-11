package node

type Role string


const (
	Follower Role = "FOLLOWER"
	Candidate Role = "CANDIDATE"
	LEADER Role = "LEADER"
)
