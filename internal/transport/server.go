package transport

import (
	"encoding/json"
	"net/http"

	"ElectraLeaderFollower/internal/node"
)

type Server struct {
	Node *node.Node
}


func New(n *node.Node) *Server {
	return &Server {
		Node: n,
	}
}

func (s *Server) Start() error {

	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.health)

	server := &http.Server{
		Addr: s.Node.Address,
		Handler: mux,
	}

	s.Node.Logger.Printf(
		"HTTP server starting on %s",
		s.Node.Address,
	)

	return server.ListenAndServe()
}

func (s *Server) health(
	w http.ResponseWriter, 
	r *http.Request,
) {
	resp := map[string]any {
		"id": s.Node.ID,
		"role": s.Node.Role,
		"address": s.Node.Address,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}