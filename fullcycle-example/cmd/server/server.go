package main

import (
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	"github.com/aferreira44/grpc-quick-start/pb"
	"github.com/aferreira44/grpc-quick-start/servers"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterMathServiceServer(grpcServer, &servers.Math{})
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":50055")

	if err != nil {
		log.Fatalf("Cannot start the grpc server: %v", err)
	}

	grpcServer.Serve(listener)
}
