package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:52000", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client := proto.NewAppServicesClient(conn)

	//request & response
	/*
		req := &proto.CalculatorRequest{X: 100, Y: 200}
		res, e := client.Add(context.Background(), req)
		if e != nil {
			log.Fatalln(e)
		}
		fmt.Println("Result from server ", res.GetResult())
	*/

	//client streaming
	fmt.Println("Average operation starts...")
	data := []int64{10, 20, 30, 40, 50}
	avgClientStream, e := client.Average(context.Background())
	if e != nil {
		log.Fatalln(e)
	}
	for _, no := range data {
		fmt.Println("Sending no : ", no)
		req := &proto.AverageRequest{No: no}
		avgClientStream.Send(req)
		time.Sleep(500 * time.Millisecond)
	}
	avgResponse, responseErr := avgClientStream.CloseAndRecv()
	if responseErr != nil {
		log.Fatalln(responseErr)
	}
	fmt.Println("Average : ", avgResponse.GetResult())
}
