package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

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

	// mysqluser := os.Getenv("DATABASE_USER")
	mysqlpass := os.Getenv("DATABASE_PASSWORD")
	mysqladdr := os.Getenv("MYSQL_PORT_3306_TCP_ADDR")

	db, err := sql.Open("mysql", "gogrpc:"+mysqlpass+"@tcp("+mysqladdr+":3306)/")
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS todosDB")
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Successfully created database..")
	}

	_, err = db.Exec("USE todosDB")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB selected successfully..")
	}

	results, err := db.Query("CREATE TABLE IF NOT EXISTS todos (id int NOT NULL AUTO_INCREMENT, title VARCHAR(40), content VARCHAR(100), PRIMARY_KEY(id));")
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Table created successfully..")
		fmt.Println(results)
	}

	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
