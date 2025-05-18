package comment_usecase

import (
	"DataTask/internal/domain/dto"
	"DataTask/internal/domain/entity"
	"context"
)

type CommentUseCase interface {
	CreateComment(ctx context.Context, commentDTO *dto.Comment, taskID int) (*entity.Comment, error)
	GetCommentsByTaskID(ctx context.Context, taskID int) ([]*dto.Comment, error)
}
