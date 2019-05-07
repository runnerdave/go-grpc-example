package main

import (
	"context"
	"fmt"
	"log"

	"github.com/runnerdave/go-grpc-example/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		A: 3,
		B: 6,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error calling the Sum service %v", err)
	}
	log.Printf("REsult from sum:%d", res.GetResult())
}
