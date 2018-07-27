package server

import (
	"log"
	"net/http"

	"github.com/Ragnar-BY/testtask_microservices/proto"
	"github.com/gorilla/mux"
)

// Server is router.
type Server struct {
	router *mux.Router
}

// NewServer creates new Server instance.
func NewServer(client proto.PlayerServiceClient) *Server {
	return &Server{newGateway(client, mux.NewRouter())}
}

// Start starts pkg with addr.
func (s *Server) Start(addr string) {
	log.Fatal(http.ListenAndServe(addr, s.router))
}
