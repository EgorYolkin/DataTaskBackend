package user_usecase

import (
	"DataTask/internal/domain/dto"
	"DataTask/internal/domain/entity"
	"DataTask/internal/repository/user_repository"

	"context"
)

type UserUseCaseImpl struct {
	repo user_repository.UserRepository
}

func NewUserUseCase(repo user_repository.UserRepository) *UserUseCaseImpl {
	return &UserUseCaseImpl{repo: repo}
}

func (uc *UserUseCaseImpl) CreateUser(ctx context.Context, user *dto.User) error {
	entityUser := entity.User{
		Email:          user.Email,
		Name:           user.Name,
		Surname:        user.Surname,
		HashedPassword: user.Password,
	}

	_, err := uc.repo.CreateUser(ctx, &entityUser)

	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUseCaseImpl) UpdateUser(ctx context.Context, user *dto.User) (*dto.User, error) {
	entityUser := entity.User{
		Email:          user.Email,
		Name:           user.Name,
		Surname:        user.Surname,
		HashedPassword: user.Password,
	}

	u, err := uc.repo.UpdateUser(ctx, &entityUser)
	if err != nil {
		return nil, err
	}

	dtoUser := dto.User{
		ID:        u.ID,
		Name:      u.Name,
		Surname:   u.Surname,
		Email:     u.Email,
		AvatarURL: u.AvatarURL,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

	return &dtoUser, nil
}

func (uc *UserUseCaseImpl) DeleteUser(ctx context.Context, id int) error {
	err := uc.repo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
