package notification_repository

import (
	"DataTask/internal/domain/entity"
	"DataTask/internal/repository/database"
	"context"
	"database/sql"
	"fmt"
)

type PostgresNotificationRepository struct {
	db *sql.DB
}

func NewPostgresNotificationRepository(db *sql.DB) *PostgresNotificationRepository {
	return &PostgresNotificationRepository{db: db}
}

func (r *PostgresNotificationRepository) CreateNotification(ctx context.Context, notification *entity.Notification) error {
	q := fmt.Sprintf(
		`
		INSERT INTO %s
			(owner_id, title, description)
		VALUES
			($1, $2, $3)
		`,
		database.NotificationTable,
	)

	_, err := r.db.ExecContext(ctx, q, notification.OwnerID, notification.Title, notification.Description)

	if err != nil {
		return fmt.Errorf("failed to create notification: %w", err)
	}

	return nil
}

func (r *PostgresNotificationRepository) GetUserNotificationsByID(ctx context.Context, ownerID int) ([]*entity.Notification, error) {
	q := fmt.Sprintf(
		`
		SELECT
			id, owner_id, title, description, is_read, created_at, updated_at
		FROM
			%s
		WHERE
			owner_id = $1
		ORDER BY
			created_at DESC
		`,
		database.NotificationTable,
	)

	rows, err := r.db.QueryContext(ctx, q, ownerID)
	if err != nil {
		return nil, fmt.Errorf("failed to query user notifications: %w", err)
	}
	defer rows.Close()

	var notifications []*entity.Notification
	for rows.Next() {
		var notification entity.Notification
		err := rows.Scan(
			&notification.ID,
			&notification.OwnerID,
			&notification.Title,
			&notification.Description,
			&notification.IsRead,
			&notification.CreatedAt,
			&notification.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan notification row: %w", err)
		}
		notifications = append(notifications, &notification)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating rows: %w", err)
	}

	return notifications, nil
}

func (r *PostgresNotificationRepository) SetNotificationIsReadByID(ctx context.Context, id int) error {
	q := fmt.Sprintf(
		`
		UPDATE %s
		SET
			is_read = TRUE,
			updated_at = NOW()
		WHERE
			id = $1
		`,
		database.NotificationTable,
	)

	result, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("failed to update notification status: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no notification found with ID %d to update", id)
	}

	return nil
}
