package node

import "sync"


type Node struct {
	ID string

	Address string

	CurrentTerm uint64


	VotedFor string

	Role Role

	Peers []string

	Logger *log.Logger

	Mu sync.RWMutex

}

func New(cfg *config.Config, logger *log.Logger) *Node {
	return &Node{
		ID: cfg.ID,
		Address: cfg.Address,
		Peers: cfg.Peers,
		Role: Follower,
		CurrentTerm: 0,
		Logger: logger,
	}
}