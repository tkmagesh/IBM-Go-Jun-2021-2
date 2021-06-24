package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
}

//request & response
func (s *Server) Add(ctx context.Context, req *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	x := req.GetX()
	y := req.GetY()
	fmt.Printf("Add request received for %d and %d\n", x, y)
	result := x + y
	response := &proto.CalculatorResponse{Result: result}
	return response, nil
}

//client streaming
func (s *Server) Average(reqStream proto.AppServices_AverageServer) error {
	var sum int64
	var count int64
	for {
		req, err := reqStream.Recv()
		if err == io.EOF {
			//prepare the response and send it
			avg := sum / count
			res := &proto.AverageResponse{Result: avg}
			reqStream.SendAndClose(res)
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		no := req.GetNo()
		sum += no
		count++
	}
	return nil
}

//server streaming
func (s *Server) GeneratePrimes(req *proto.PrimeRequest, serverStream proto.AppServices_GeneratePrimesServer) error {
	rangeStart := req.GetRangeStart()
	rangeEnd := req.GetRangeEnd()
	for no := rangeStart; no <= rangeEnd; no++ {
		if isPrime(no) {
			fmt.Println("Sending prime no : ", no)
			res := &proto.PrimeResponse{No: no}
			serverStream.Send(res)
			time.Sleep(500 * time.Millisecond)
		}
	}
	return nil
}

func isPrime(no int64) bool {
	for i := int64(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func (s *Server) Greet(reqResStream proto.AppServices_GreetServer) error {
	fmt.Println("Greet everyone function is invoked")
	for {
		req, err := reqResStream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalln(err)
		}
		greeting := req.GetGreeting()
		firstName := greeting.GetFirstName()
		lastName := greeting.GetLastName()
		greetMsg := fmt.Sprintf("Hi %s %s, Have a nice day!", firstName, lastName)
		res := &proto.GreetingResponse{Message: greetMsg}
		sendErr := reqResStream.Send(res)
		if sendErr != nil {
			log.Fatalln(sendErr)
			return sendErr
		}
	}
	return nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:52000")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServicesServer(grpcServer, &Server{})
	e := grpcServer.Serve(listener)
	if e != nil {
		log.Fatalln(e)
	}
}
