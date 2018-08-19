package main

import (
	"context"
	"errors"
	"fmt"
	"net"

	"app/server/cmd"
	pb "app/todo"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type todoServer struct {
	todos []*pb.Todo
}

func (s *todoServer) Search(ctx context.Context, searchReq *pb.SearchReq) (*pb.SearchRes, error) {
	if len(s.todos) == 0 {
		return nil, errors.New("404 no todos")
	}
	res := &pb.SearchRes{
		Todos: s.todos,
	}
	return res, nil
}

func (s *todoServer) Save(ctx context.Context, param *pb.TodoToSave) (*pb.Ack, error) {
	todo := &pb.Todo{
		Content:   param.Content,
		Completed: false,
		Owner:     param.Owner,
	}

	s.todos = append(s.todos, todo)

	return &pb.Ack{Code: 200, Message: "OK"}, nil
}

func newServer() *todoServer {
	return &todoServer{
		todos: make([]*pb.Todo, 0),
	}
}

func main() {
	cmd.Execute()
	port := viper.GetInt("port")
	host := viper.GetString("host")

	conn, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(fmt.Sprintf("err %s", err))
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTodoServiceServer(grpcServer, newServer())
	fmt.Println("Listening on port", port)
	grpcServer.Serve(conn)

}
