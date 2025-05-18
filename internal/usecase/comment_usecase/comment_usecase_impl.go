package comment_usecase

import (
	"DataTask/internal/domain/dto"
	"DataTask/internal/domain/entity"
	"DataTask/internal/repository/comment_repository"
	"context"
	"fmt"
)

type CommentUseCaseImpl struct {
	repo comment_repository.CommentRepository
}

func NewCommentUseCase(repo comment_repository.CommentRepository) *CommentUseCaseImpl {
	return &CommentUseCaseImpl{repo: repo}
}

func (uc *CommentUseCaseImpl) CreateComment(ctx context.Context, commentDTO *dto.Comment, taskID int) (*entity.Comment, error) {
	if commentDTO.Author == nil || commentDTO.Author.ID == 0 {
		return nil, fmt.Errorf("author ID is required")
	}

	commentEntity := &entity.Comment{
		Author: &entity.User{
			ID: commentDTO.Author.ID,
		},
		Text: commentDTO.Text,
	}

	createdCommentEntity, err := uc.repo.CreateCommentForTask(ctx, commentEntity, taskID)
	if err != nil {
		return nil, fmt.Errorf("create comment in repository: %w", err)
	}
	return createdCommentEntity, nil
}

func (uc *CommentUseCaseImpl) GetCommentsByTaskID(ctx context.Context, taskID int) ([]*dto.Comment, error) {
	comments, err := uc.repo.GetCommentsByTaskID(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("get comments from repository: %w", err)
	}

	commentDTOs := make([]*dto.Comment, len(comments))
	for i, comment := range comments {
		commentDTOs[i] = &dto.Comment{
			ID: comment.ID,
			Author: &dto.User{
				ID:        comment.Author.ID,
				Name:      comment.Author.Name,
				Surname:   comment.Author.Surname,
				Email:     comment.Author.Email,
				AvatarURL: comment.Author.AvatarURL,
				CreatedAt: comment.Author.CreatedAt,
				UpdatedAt: comment.Author.UpdatedAt,
			},
			Text:      comment.Text,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		}
	}
	return commentDTOs, nil
}
