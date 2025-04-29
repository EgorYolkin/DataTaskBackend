package task_usecase

import (
	"DataTask/internal/domain/dto"
	"context"
)

type TaskUseCase interface {
	CreateTask(ctx context.Context, task *dto.Task) (*dto.Task, error)
	GetTaskByID(ctx context.Context, id int) (*dto.Task, error)
	UpdateTask(ctx context.Context, task *dto.Task) (*dto.Task, error)
	DeleteTask(ctx context.Context, id int) error
	GetTasksByKanbanID(ctx context.Context, kanbanID int) ([]*dto.Task, error)
	GetTasksByUserID(ctx context.Context, userID int) ([]*dto.Task, error)
	AssignUserToTask(ctx context.Context, taskID int, userID int) error
	GetTasksByProjectID(ctx context.Context, projectID int) ([]*dto.Task, error)
}
