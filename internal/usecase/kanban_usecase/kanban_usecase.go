package kanban_usecase

import (
	"DataTask/internal/domain/dto"
	"context"
)

type KanbanUseCase interface {
	CreateKanban(ctx context.Context, kanban *dto.Kanban) (*dto.Kanban, error)
	GetKanbanByID(ctx context.Context, id int) (*dto.Kanban, error)
	GetKanbansByProjectID(ctx context.Context, projectID int) ([]*dto.Kanban, error)
	UpdateKanban(ctx context.Context, kanban *dto.Kanban) (*dto.Kanban, error)
	DeleteKanban(ctx context.Context, id int) error
	GetAllKanbans(ctx context.Context) ([]*dto.Kanban, error)
}
