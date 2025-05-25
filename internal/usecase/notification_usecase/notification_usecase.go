package notification_usecase

import (
	"DataTask/internal/domain/dto"
	"context"
)

type NotificationUseCase interface {
	CreateNotification(ctx context.Context, notification *dto.Notification) error
	GetUserNotificationsByID(ctx context.Context, id int) ([]*dto.Notification, error)
	SetNotificationIsReadByID(ctx context.Context, id int) error
}
