package task_repository

import (
	"DataTask/internal/domain/entity"
	"context"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task *entity.Task) (*entity.Task, error)
	GetTaskByID(ctx context.Context, id int) (*entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) (*entity.Task, error)
	DeleteTask(ctx context.Context, id int) error
	GetTasksByKanbanID(ctx context.Context, kanbanID int) ([]*entity.Task, error)
	GetTasksByUserID(ctx context.Context, userID int) ([]*entity.Task, error)
	AssignUserToTask(ctx context.Context, taskID int, userID int) error
	GetTasksByProjectID(ctx context.Context, projectID int) ([]*entity.Task, error)
}
