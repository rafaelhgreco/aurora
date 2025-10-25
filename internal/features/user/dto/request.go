package dto

// @Description Representa o request de criação de um novo usuário
type CreateUserRequest struct {
	// @Seu nome
	Name        string   `json:"name" binding:"required"`
	// @SeuEmail@email.com
	Email       string   `json:"email" binding:"required,email"`
	// @SuaSenha123@
	Password    string   `json:"password" binding:"required,min=8"`
	// @ADMIN
	// @COLLABORATOR
	// @COMMON
	Type        string   `json:"type" binding:"oneof=COMMON COLLABORATOR ADMIN ''"`
	Permissions []string `json:"permissions"`
	TeamID      string   `json:"teamId"`
}
// @Description Representa o request de atualização de um usuário
type UpdateUserRequest struct {
	// @Seu novo nome
	Name  *string `json:"name,omitempty"`
	// @SeuNovoEmail@email.com
	Email *string `json:"email,omitempty" binding:"omitempty,email"`
}

// @Description Representa o request de exclusão de um usuário
type DeleteUserRequest struct {
	// @ID do usuário
	ID string `json:"id" binding:"required"`
}
