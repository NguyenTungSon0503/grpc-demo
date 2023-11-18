package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/NguyenTungSon0503/grpc-demo/proto/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "Hello " + req.Name,
		Age:     "Your Age is " + req.Age,
	}, nil
}

type calculateServer struct {
	pb.UnimplementedCalculateServer
}

func (s *calculateServer) Summary(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	num1 := req.GetNum1()
	num2 := req.GetNum2()
	sum := num1 + num2
	fmt.Println(num1, num2)
	return &pb.SumResponse{Sum: sum}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	greeter := &greeterServer{}
	calculate := &calculateServer{}

	pb.RegisterGreeterServer(grpcServer, greeter)
	pb.RegisterCalculateServer(grpcServer, calculate)

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
