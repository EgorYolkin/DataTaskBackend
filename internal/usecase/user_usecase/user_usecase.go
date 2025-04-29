package user_usecase

import (
	"DataTask/internal/domain/dto"
	"DataTask/internal/domain/entity"
	"context"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, user *dto.User) error
	UpdateUser(ctx context.Context, user *dto.User) (*dto.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetUserEntityByEmail(ctx context.Context, email string) (*entity.User, error)
}
