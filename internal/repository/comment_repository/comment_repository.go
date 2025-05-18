package comment_repository

import (
	"DataTask/internal/domain/entity"
	"context"
)

type CommentRepository interface {
	CreateCommentForTask(ctx context.Context, comment *entity.Comment, TaskID int) (*entity.Comment, error)
	GetCommentsByTaskID(ctx context.Context, taskID int) ([]*entity.Comment, error)
}
