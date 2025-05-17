package dto

import "time"

type Project struct {
	ID              int    `json:"id"`
	OwnerID         int    `json:"owner_id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Color           string `json:"color"`
	ParentProjectID *int   `json:"parent_project_id"`
	//  Убрали CreatedAt и UpdatedAt из запроса на создание
	CreatedAt time.Time `json:"created_at,omitempty"` //  omitempty, чтобы не возвращать null
	UpdatedAt time.Time `json:"updated_at,omitempty"` //  omitempty, чтобы не возвращать null
}

type ProjectUser struct {
	ProjectID       int        `json:"project_id"`
	UserID          int        `json:"user_id"`
	Permission      string     `json:"permission"`
	InvitedByUserID *int       `json:"invited_by_user_id"`
	InvitedAt       time.Time  `json:"invited_at"`
	JoinedAt        *time.Time `json:"joined_at"`
}

type ProjectUserInvite struct {
	ProjectID  int    `json:"project_id" binding:"required"`
	UserID     int    `json:"user_id"`
	Permission string `json:"permission"`
}
