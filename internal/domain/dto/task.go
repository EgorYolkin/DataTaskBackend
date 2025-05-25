package dto

import "time"

type Task struct {
	ID          int       `json:"id"`
	OwnerID     int       `json:"owner_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	KanbanID    int       `json:"kanban_id"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
