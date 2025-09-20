package dto

import "time"

type UserResponse struct {
	ID        string       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Type 	string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type AdminUserResponse struct {
	UserResponse
	Permissions []string `json:"permissions"`
}

type CollaboratorUserResponse struct {
	UserResponse
	TeamID string   `json:"team_id"`
	Projects []string `json:"projects"`
}

type CommonUserResponse = UserResponse