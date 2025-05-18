package entity

import "time"

type Comment struct {
	ID        int       `json:"id"`
	Author    *User     `json:"author"`
	Text      string    `json:"text"`
	TaskID    int       `json:"task_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
