package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"../api"
)

// starts a gRPC server and waits for connection
func main() {
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a server instance
	s := api.Server{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	api.RegisterPingServer(grpcServer, &s)

	// start the server
	serve_err := grpcServer.Serve(lis)
	if serve_err != nil {
		log.Fatalf("failed to serve: %s", serve_err)
	}
}
