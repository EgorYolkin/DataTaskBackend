package kanban_usecase

import (
	"DataTask/internal/domain/dto"
	"DataTask/internal/domain/entity"
	"DataTask/internal/repository/kanban_repository"
	"context"
)

type KanbanUseCaseImpl struct {
	repo kanban_repository.KanbanRepository
}

func NewKanbanUseCase(repo kanban_repository.KanbanRepository) *KanbanUseCaseImpl {
	return &KanbanUseCaseImpl{repo: repo}
}

func (uc *KanbanUseCaseImpl) CreateKanban(ctx context.Context, kanban *dto.Kanban) (*dto.Kanban, error) {
	entityKanban := &entity.Kanban{
		Name:      kanban.Name,
		ProjectID: kanban.ProjectID,
	}

	createdKanban, err := uc.repo.CreateKanban(ctx, entityKanban)
	if err != nil {
		return nil, err
	}

	dtoKanban := &dto.Kanban{
		ID:        createdKanban.ID,
		Name:      createdKanban.Name,
		ProjectID: createdKanban.ProjectID,
		CreatedAt: createdKanban.CreatedAt,
		UpdatedAt: createdKanban.UpdatedAt,
	}

	return dtoKanban, nil
}

func (uc *KanbanUseCaseImpl) GetKanbanByID(ctx context.Context, id int) (*dto.Kanban, error) {
	kanban, err := uc.repo.GetKanbanByID(ctx, id)
	if err != nil {
		return nil, err
	}

	dtoKanban := &dto.Kanban{
		ID:        kanban.ID,
		Name:      kanban.Name,
		CreatedAt: kanban.CreatedAt,
		UpdatedAt: kanban.UpdatedAt,
	}

	return dtoKanban, nil
}

func (uc *KanbanUseCaseImpl) GetKanbansByProjectID(ctx context.Context, projectID int) ([]*dto.Kanban, error) {
	kanbans, err := uc.repo.GetKanbansByProjectID(ctx, projectID)
	if err != nil {
		return nil, err
	}

	var dtoKanbans []*dto.Kanban

	for _, kanban := range kanbans {
		dtoKanban := &dto.Kanban{
			ID:        kanban.ID,
			Name:      kanban.Name,
			CreatedAt: kanban.CreatedAt,
			UpdatedAt: kanban.UpdatedAt,
		}

		dtoKanbans = append(dtoKanbans, dtoKanban)
	}

	return dtoKanbans, nil
}

func (uc *KanbanUseCaseImpl) UpdateKanban(ctx context.Context, kanban *dto.Kanban) (*dto.Kanban, error) {
	entityKanban := &entity.Kanban{
		ID:   kanban.ID,
		Name: kanban.Name,
	}

	updatedKanban, err := uc.repo.UpdateKanban(ctx, entityKanban)
	if err != nil {
		return nil, err
	}

	dtoKanban := &dto.Kanban{
		ID:        updatedKanban.ID,
		Name:      updatedKanban.Name,
		CreatedAt: updatedKanban.CreatedAt,
		UpdatedAt: updatedKanban.UpdatedAt,
	}

	return dtoKanban, nil
}

func (uc *KanbanUseCaseImpl) DeleteKanban(ctx context.Context, id int) error {
	err := uc.repo.DeleteKanban(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *KanbanUseCaseImpl) GetAllKanbans(ctx context.Context) ([]*dto.Kanban, error) {
	kanbans, err := uc.repo.GetAllKanbans(ctx)
	if err != nil {
		return nil, err
	}

	var dtoKanbans []*dto.Kanban
	for _, k := range kanbans {
		dtoKanbans = append(dtoKanbans, &dto.Kanban{
			ID:        k.ID,
			Name:      k.Name,
			CreatedAt: k.CreatedAt,
			UpdatedAt: k.UpdatedAt,
		})
	}

	return dtoKanbans, nil
}
