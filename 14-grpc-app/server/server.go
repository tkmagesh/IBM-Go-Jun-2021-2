package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
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
