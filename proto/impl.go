package proto

import "context"

type IService interface {
	Create(context.Context, CreateTaskRequest) (CreateTaskResponse, error)
	GetTask(context.Context, GetTaskRequest) (GetTaskResponse, error)
	GetTasks(context.Context, GetTasksRequest) (GetTasksResponse, error)
}
