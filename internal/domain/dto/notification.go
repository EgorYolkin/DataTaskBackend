package dto

import "time"

type Notification struct {
	ID      int `json:"id"`
	OwnerID int `json:"owner_id"`

	Title       string `json:"title"`
	Description string `json:"description"`

	IsRead bool `json:"is_read"`

	CreatedAt time.Time `json:"created_at,omitempty"` //  omitempty, чтобы не возвращать null
	UpdatedAt time.Time `json:"updated_at,omitempty"` //  omitempty, чтобы не возвращать null
}
