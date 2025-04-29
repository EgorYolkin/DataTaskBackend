package user_usecase

import (
	"DataTask/internal/domain/dto"
	"DataTask/internal/domain/entity"
	"DataTask/internal/repository/user_repository"
	"DataTask/pkg/hashing"

	"context"
)

type UserUseCaseImpl struct {
	repo user_repository.UserRepository
}

func NewUserUseCase(repo user_repository.UserRepository) *UserUseCaseImpl {
	return &UserUseCaseImpl{repo: repo}
}

func (uc *UserUseCaseImpl) CreateUser(ctx context.Context, user *dto.User) error {
	hashOptions := hashing.DefaultHashOptions
	hashOptions.Value = user.Password

	hashedPassword, err := hashing.Hash(hashOptions)
	if err != nil {
		return err
	}

	entityUser := entity.User{
		Email:          user.Email,
		Name:           user.Name,
		Surname:        user.Surname,
		HashedPassword: hashedPassword,
	}

	_, err = uc.repo.CreateUser(ctx, &entityUser)

	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUseCaseImpl) UpdateUser(ctx context.Context, user *dto.User) (*dto.User, error) {
	entityUser := entity.User{
		Email:   user.Email,
		Name:    user.Name,
		Surname: user.Surname,
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

func (uc *UserUseCaseImpl) GetUserEntityByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := uc.repo.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
