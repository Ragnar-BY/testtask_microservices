package main

import (
	"log"

	"github.com/Ragnar-BY/testtask_microservices/apigateway/server"
	"github.com/Ragnar-BY/testtask_microservices/proto"
	"google.golang.org/grpc"
)

// Connect creates grpc connection and starts server.
func Connect(opts settings) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(opts.PlayerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	c := proto.NewPlayerServiceClient(conn)

	srv := server.NewServer(c)
	srv.Start(opts.ServerAddress)
}
func main() {
	opts := new(settings)
	err := opts.Parse()
	if err != nil {
		log.Fatalf("Cannot parse settings: %v", err)
	}
	Connect(*opts)
}
