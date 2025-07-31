package main

import (
	"log"
	"net"

	pb "github.com/aman-zulfiqar/go-grpcx/proto"
	"google.golang.org/grpc"
)

const (
	// Port is the port on which the server will run.
	Port = ":8080"
)

// helloServer implements the GreetServiceServer interface defined in the proto file.
type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	listener, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

	// Create a new gRPC server and register the helloServer.
	// This server will handle incoming gRPC requests.
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Panicf("Server is listening on port %s", listener.Addr().String())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)

	}
}
