package dto

type CreateUserRequest struct {
	Name        string   `json:"name" binding:"required"`
	Email       string   `json:"email" binding:"required,email"`
	Password    string   `json:"password" binding:"required,min=8"`
	Type        string   `json:"type" binding:"required,oneof=COMMON COLLABORATOR ADMIN"`
	Permissions []string `json:"permissions,omitempty"`
	TeamID      string   `json:"teamId,omitempty"`
}

type UpdateUserRequest struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty" binding:"omitempty,email"`
}

type DeleteUserRequest struct {
	ID string `json:"id" binding:"required"`
}