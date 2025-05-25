package di

import (
	"DataTask/internal/repository/notification_repository"
	"DataTask/internal/usecase/notification_usecase"
	"database/sql"
)

func InitializeNotificationUseCase(db *sql.DB) notification_usecase.NotificationUseCase {
	repo := notification_repository.NewPostgresNotificationRepository(db)
	useCase := notification_usecase.NewNotificationUseCase(repo)

	return useCase
}
