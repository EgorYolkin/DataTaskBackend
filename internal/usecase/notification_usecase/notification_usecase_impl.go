package notification_usecase

import (
	"DataTask/internal/domain/dto"
	"DataTask/internal/domain/entity" // Импортируем entity
	"DataTask/internal/repository/notification_repository"
	"context"
	"fmt"
)

type NotificationUseCaseImpl struct {
	repo notification_repository.NotificationRepository
}

func NewNotificationUseCase(repo notification_repository.NotificationRepository) *NotificationUseCaseImpl {
	return &NotificationUseCaseImpl{repo: repo}
}

// CreateNotification преобразует DTO в сущность и вызывает метод репозитория.
func (uc *NotificationUseCaseImpl) CreateNotification(ctx context.Context, notificationDTO *dto.Notification) error {
	// Преобразуем DTO в Entity
	notificationEntity := &entity.Notification{
		OwnerID:     notificationDTO.OwnerID,
		Title:       notificationDTO.Title,
		Description: notificationDTO.Description,
		// ID, IsRead, CreatedAt, UpdatedAt обычно устанавливаются базой данных или репозиторием
	}

	err := uc.repo.CreateNotification(ctx, notificationEntity)
	if err != nil {
		return fmt.Errorf("usecase: failed to create notification: %w", err)
	}

	return nil
}

// GetUserNotificationsByID получает уведомления от репозитория и преобразует их в DTO.
func (uc *NotificationUseCaseImpl) GetUserNotificationsByID(ctx context.Context, ownerID int) ([]*dto.Notification, error) {
	notificationEntities, err := uc.repo.GetUserNotificationsByID(ctx, ownerID)
	if err != nil {
		return nil, fmt.Errorf("usecase: failed to get user notifications: %w", err)
	}

	var notificationDTOs []*dto.Notification
	for _, ent := range notificationEntities {
		// Преобразуем Entity в DTO
		dto := &dto.Notification{
			ID:          ent.ID,
			OwnerID:     ent.OwnerID,
			Title:       ent.Title,
			Description: ent.Description,
			IsRead:      ent.IsRead,
			CreatedAt:   ent.CreatedAt,
			UpdatedAt:   ent.UpdatedAt,
		}
		notificationDTOs = append(notificationDTOs, dto)
	}

	return notificationDTOs, nil
}

// SetNotificationIsReadByID вызывает метод репозитория для установки статуса прочитанного.
func (uc *NotificationUseCaseImpl) SetNotificationIsReadByID(ctx context.Context, id int) error {
	err := uc.repo.SetNotificationIsReadByID(ctx, id)
	if err != nil {
		return fmt.Errorf("usecase: failed to set notification as read: %w", err)
	}
	return nil
}
