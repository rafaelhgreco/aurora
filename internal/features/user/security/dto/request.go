package dto

type LoginRequest struct {
	IDToken string `json:"id_token" binding:"required"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type ChangePasswordRequest struct {
	NewPassword     string `json:"newPassword" binding:"required,min=8"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=NewPassword"`
}
