package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) Add(ctx context.Context, req *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	x := req.GetX()
	y := req.GetY()
	fmt.Printf("Add request received for %d and %d\n", x, y)
	result := x + y
	response := &proto.CalculatorResponse{Result: result}
	return response, nil
}

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
