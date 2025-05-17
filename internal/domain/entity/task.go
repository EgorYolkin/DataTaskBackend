package entity

import "time"

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	KanbanID    int       `json:"kanban_id"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
