package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/waveywaves/grpc-example/proto"
	"google.golang.org/grpc"
)

const (
	port = ":3000"
)

type server struct{}

func (s *server) Create(ctx context.Context, todo *pb.Todo) (*pb.Todo, error) {
	fmt.Println(todo)
	return todo, nil
}

func (s *server) Delete(ctx context.Context, todo *pb.TodoRequestId) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *server) List(ctx context.Context, empty *pb.Empty) (*pb.TodoList, error) {
	return &pb.TodoList{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
