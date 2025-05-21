package entity

import "time"

type Project struct {
	ID              int       `json:"id"`
	OwnerID         int       `json:"owner_id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Color           string    `json:"color"`
	ParentProjectID *int      `json:"parent_project_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type ProjectUser struct {
	ProjectID       int        `json:"project_id"`
	UserID          int        `json:"user_id"`
	UserEmail       string     `json:"user_email"`
	Permission      string     `json:"permission"`
	InvitedByUserID *int       `json:"invited_by_user_id"`
	InvitedAt       time.Time  `json:"invited_at"`
	JoinedAt        *time.Time `json:"joined_at"`
}
