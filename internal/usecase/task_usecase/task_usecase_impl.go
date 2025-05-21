package task_usecase

import (
	"DataTask/internal/domain/dto"
	"DataTask/internal/domain/entity"
	"DataTask/internal/repository/task_repository"
	"context"
)

type TaskUseCaseImpl struct {
	repo task_repository.TaskRepository
}

func NewTaskUseCase(repo task_repository.TaskRepository) *TaskUseCaseImpl {
	return &TaskUseCaseImpl{repo: repo}
}

func (uc *TaskUseCaseImpl) CreateTask(ctx context.Context, task *dto.Task) (*dto.Task, error) {
	entityTask := &entity.Task{
		Title:       task.Title,
		KanbanID:    task.KanbanID,
		Description: task.Description,
		IsCompleted: task.IsCompleted,
	}

	createdTask, err := uc.repo.CreateTask(ctx, entityTask)
	if err != nil {
		return nil, err
	}

	dtoTask := &dto.Task{
		ID:          createdTask.ID,
		KanbanID:    createdTask.KanbanID,
		Title:       createdTask.Title,
		Description: createdTask.Description,
		IsCompleted: createdTask.IsCompleted,
		CreatedAt:   createdTask.CreatedAt,
		UpdatedAt:   createdTask.UpdatedAt,
	}

	return dtoTask, nil
}

func (uc *TaskUseCaseImpl) GetTaskByID(ctx context.Context, id int) (*dto.Task, error) {
	task, err := uc.repo.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}

	dtoTask := &dto.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		IsCompleted: task.IsCompleted,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}

	return dtoTask, nil
}

func (uc *TaskUseCaseImpl) UpdateTask(ctx context.Context, task *dto.Task) (*dto.Task, error) {
	entityTask := &entity.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		IsCompleted: task.IsCompleted,
	}

	updatedTask, err := uc.repo.UpdateTask(ctx, entityTask)
	if err != nil {
		return nil, err
	}

	dtoTask := &dto.Task{
		ID:          updatedTask.ID,
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		IsCompleted: updatedTask.IsCompleted,
		CreatedAt:   updatedTask.CreatedAt,
		UpdatedAt:   updatedTask.UpdatedAt,
	}

	return dtoTask, nil
}

func (uc *TaskUseCaseImpl) DeleteTask(ctx context.Context, id int) error {
	err := uc.repo.DeleteTask(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *TaskUseCaseImpl) GetTasksByKanbanID(ctx context.Context, kanbanID int) ([]*dto.Task, error) {
	tasks, err := uc.repo.GetTasksByKanbanID(ctx, kanbanID)
	if err != nil {
		return nil, err
	}

	var dtoTasks []*dto.Task
	for _, t := range tasks {
		dtoTasks = append(dtoTasks, &dto.Task{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			IsCompleted: t.IsCompleted,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
		})
	}

	return dtoTasks, nil
}

func (uc *TaskUseCaseImpl) GetTasksByUserID(ctx context.Context, userID int) ([]*dto.Task, error) {
	tasks, err := uc.repo.GetTasksByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var dtoTasks []*dto.Task
	for _, t := range tasks {
		dtoTasks = append(dtoTasks, &dto.Task{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			IsCompleted: t.IsCompleted,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
		})
	}

	return dtoTasks, nil
}

func (uc *TaskUseCaseImpl) AssignUserToTask(ctx context.Context, taskID int, userID int) error {
	err := uc.repo.AssignUserToTask(ctx, taskID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *TaskUseCaseImpl) GetTasksByProjectID(ctx context.Context, projectID int) ([]*dto.Task, error) {
	tasks, err := uc.repo.GetTasksByProjectID(ctx, projectID)
	if err != nil {
		return nil, err
	}

	var dtoTasks []*dto.Task
	for _, t := range tasks {
		dtoTasks = append(dtoTasks, &dto.Task{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			IsCompleted: t.IsCompleted,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
		})
	}

	return dtoTasks, nil
}
