package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"microservice_project/proto-gen/dice"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// TutorialService is a struct that implements the server interface

type TutorialService struct {
	dice.UnimplementedTutorialServer
}

func (s *TutorialService) RollDice(ctx context.Context, req *dice.RollDiceRequest) (*dice.RollDiceResponse, error) {
	var res dice.RollDiceResponse

	if req.Num < 0 {
		return nil, grpc.Errorf(codes.InvalidArgument, "number should be positive")
	}

	for i := 0; i < int(req.Num); i++ {
		res.Dice = append(res.Dice, int32(rand.Intn(100)))
		fmt.Println(rand.Intn(100))
	}

	return &res, nil
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
