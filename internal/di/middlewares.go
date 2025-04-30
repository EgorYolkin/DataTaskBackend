package di

import (
	"DataTask/internal/controller/rest/middleware/auth_middleware"
	"DataTask/internal/repository/user_repository"
	"DataTask/internal/usecase/user_usecase"
	"database/sql"
)

func InitializeAuthMiddleware(db *sql.DB, jwtSecretKey string) *auth_middleware.AuthMiddleware {
	repo := user_repository.NewPostgresUserRepository(db)
	useCase := user_usecase.NewUserUseCase(repo)
	middleware := auth_middleware.NewAuthMiddleware(useCase, jwtSecretKey)
	return middleware
}
