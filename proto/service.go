package proto

import (
	"context"

	"github.com/oliver7100/task-service/database"
)

type service struct {
	UnimplementedTaskServiceServer
	Conn *database.Connection
}

func (s *service) Create(ctx context.Context, req *CreateTaskRequest) (*CreateTaskResponse, error) {
	return nil, nil
}

func (s *service) GetTask(ctx context.Context, req *GetTaskRequest) (*GetTaskResponse, error) {
	return nil, nil
}

func (s *service) GetTasks(ctx context.Context, req *GetTasksRequest) (*GetTasksResponse, error) {
	return nil, nil
}

func NewTaskService(conn *database.Connection) *service {
	return &service{
		Conn: conn,
	}
}
