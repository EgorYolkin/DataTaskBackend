package kanban_repository

import (
	"DataTask/internal/domain/entity"
	"context"
)

type KanbanRepository interface {
	CreateKanban(ctx context.Context, kanban *entity.Kanban) (*entity.Kanban, error)
	GetKanbanByID(ctx context.Context, id int) (*entity.Kanban, error)
	GetKanbansByProjectID(ctx context.Context, projectID int) ([]*entity.Kanban, error)
	UpdateKanban(ctx context.Context, kanban *entity.Kanban) (*entity.Kanban, error)
	DeleteKanban(ctx context.Context, id int) error
	GetAllKanbans(ctx context.Context) ([]*entity.Kanban, error) // Added GetAll
}
