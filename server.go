package main

import (
	"context"
	"fmt"
	"log"
	"microservice_project/proto-gen/dice"
	"net"

	"google.golang.org/grpc"
)

// TutorialService is a struct that implements the server interface

type TutorialService struct {
	dice.UnimplementedTutorialServer
}

func (s *TutorialService) RollDice(ctx context.Context, req *dice.RollDiceRequest) (*dice.RollDiceResponse, error) {
	fmt.Printf("req -----> %#+v", req)
	return nil, nil
}

func main() {
	println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	dice.RegisterTutorialServer(s, &TutorialService{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}
