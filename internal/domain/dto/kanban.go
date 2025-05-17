package dto

import "time"

type Kanban struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	ProjectID int       `json:"project_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}
