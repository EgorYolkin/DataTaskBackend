package entity

import "time"

type User struct {
	ID int `json:"id,omitempty"`

	Name      string  `json:"name,omitempty"`
	Surname   string  `json:"surname,omitempty"`
	Email     string  `json:"email"`
	AvatarURL *string `json:"avatar_url,omitempty"`

	HashedPassword string `json:"-"`
	Salt           string `json:"-"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
