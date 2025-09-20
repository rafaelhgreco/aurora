package dto

import userDTO "aurora.com/aurora-backend/internal/features/user/dto"

type LoginResponse struct {
	User         *userDTO.UserResponse `json:"user"`
	AccessToken  string                `json:"accessToken"`
	RefreshToken string                `json:"refreshToken"`
}