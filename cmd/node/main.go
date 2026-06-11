package main

import (
	"os"

	"github.com/himanshu/ElectraLeaderFollower/internal/config"
	"github.com/himanshu/ElectraLeaderFollower/internal/logger"
	"github.com/himanshu/ElectraLeaderFollower/internal/node"
	"github.com/himanshu/ElectraLeaderFollower/internal/transport"
)



func main() {
	if len(os.Args) != 2 {
		panic(
			"usage: go run main.go config.json",
		)
	}
	cfg := config.Load(os.Args[1])

	log := logger.New()

	n := node.New(cfg, log)

	server := transport.New(n)

	if err := server.Start(); err != nil {
		panic(err)
	}
}
