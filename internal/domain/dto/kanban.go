package dto

import "time"

type Kanban struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	//  Убрали CreatedAt и UpdatedAt из запроса на создание
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}
