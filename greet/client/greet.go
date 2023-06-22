package main

import (
	"context"
	"log"

	pb "github.com/suhelzende/go-grpc/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("Greeting ...")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Suhel",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Greeting: %s", res.Result)
}
