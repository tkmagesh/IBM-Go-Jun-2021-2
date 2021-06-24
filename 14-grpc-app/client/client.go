package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:52000", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client := proto.NewAppServicesClient(conn)
	req := &proto.CalculatorRequest{X: 100, Y: 200}
	res, e := client.Add(context.Background(), req)
	if e != nil {
		log.Fatalln(e)
	}
	fmt.Println("Result from server ", res.GetResult())
}
