package user_repository

import (
	"DataTask/internal/domain/entity"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}
