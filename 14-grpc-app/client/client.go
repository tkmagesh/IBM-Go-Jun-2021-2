package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
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
	/*
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
	*/

	//server streaming
	/*
		fmt.Println("Prime numbers between 7 and 63")
		req := &proto.PrimeRequest{RangeStart: 7, RangeEnd: 63}
		primeNoStream, e := client.GeneratePrimes(context.Background(), req)
		if e != nil {
			log.Fatalln(e)
		}
		for {
			res, er := primeNoStream.Recv()
			if er == io.EOF {
				fmt.Println("All the prime numbers are received")
				break
			}
			if er != nil {
				log.Fatalln(er)
			}
			fmt.Println("Prime No received : ", res.GetNo())
		}
	*/

	//bidirectional streaming
	names := []proto.Greeting{
		proto.Greeting{FirstName: "Magesh", LastName: "Kuppan"},
		proto.Greeting{FirstName: "Ramesh", LastName: "Jayaraman"},
		proto.Greeting{FirstName: "Suresh", LastName: "Kannan"},
		proto.Greeting{FirstName: "Guru", LastName: "Raghav"},
	}
	stream, err := client.Greet(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	//send the messages concurrently
	go func() {
		for _, greetingName := range names {
			req := &proto.GreetingRequest{
				Greeting: &greetingName,
			}
			fmt.Println("Sending req : ", greetingName.FirstName, greetingName.LastName)
			stream.Send(req)
			time.Sleep(500 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	closeCh := make(chan string)
	//recieve the responses concurrently
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println("Server response : ", res.GetMessage())
		}
		closeCh <- "done"
	}()

	<-closeCh
}
