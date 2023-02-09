package main

import (
	"net"

	"github.com/oliver7100/task-service/database"
	"github.com/oliver7100/task-service/proto"
	"google.golang.org/grpc"
)

func main() {
	dbConnection, err := database.NewDatabaseConnection(
		"root:root@tcp(127.0.0.1:3306)/db_task_service?charset=utf8mb4&parseTime=True&loc=Local",
	)

	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	proto.RegisterTaskServiceServer(
		s,
		proto.NewTaskService(
			dbConnection,
		),
	)

	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
