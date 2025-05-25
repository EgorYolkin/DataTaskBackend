package notification_repository

import (
	"DataTask/internal/domain/entity"
	"context"
)

type NotificationRepository interface {
	CreateNotification(ctx context.Context, notification *entity.Notification) error
	GetUserNotificationsByID(ctx context.Context, id int) ([]*entity.Notification, error)
	SetNotificationIsReadByID(ctx context.Context, id int) (error)
}