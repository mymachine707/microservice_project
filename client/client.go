package main

import (
	"context"
	"log"
	"time"

	"microservice_project/proto-gen/dice"

	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := dice.NewTutorialClient(conn)

	for i := 0; i < 100; i++ {
		res, err := c.Ping(context.Background(), &dice.Empty{})

		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}
		log.Printf("Response from server: %s", res)
		time.Sleep(time.Second * 1)
	}

}
