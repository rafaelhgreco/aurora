package dto

import "time"

// UserResponse é o DTO para expor os dados de um usuário. Note que não expõe a senha.
type UserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}