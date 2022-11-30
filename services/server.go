package main

import (
	"log"
	"microservice_project/proto-gen/dice"
	dService "microservice_project/services/service/dice"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	dice.RegisterTutorialServer(s, &dService.TutorialService{})

	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}
