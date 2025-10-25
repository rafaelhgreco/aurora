package controller

import (
	"net/http"

	"aurora.com/aurora-backend/internal/features/user/dto"
	"github.com/gin-gonic/gin"

	"aurora.com/aurora-backend/internal/features/user/mapper"
	securityDTO "aurora.com/aurora-backend/internal/features/user/security/dto"
	userAuthUseCase "aurora.com/aurora-backend/internal/features/user/security/use-case"
	userUseCase "aurora.com/aurora-backend/internal/features/user/use-case"
)

type UserController struct {
	createUser         *userUseCase.CreateUserUseCase
	updateUser         *userUseCase.UpdateUserUseCase
	getUserById        *userUseCase.GetUserByIDUseCase
	deleteUser         *userUseCase.DeleteUserUseCase
	loginUser          *userAuthUseCase.LoginUserUseCase
	changePasswordUser *userAuthUseCase.ChangePasswordUseCase
}

func NewUserController(createUser *userUseCase.CreateUserUseCase, updateUser *userUseCase.UpdateUserUseCase, getUserById *userUseCase.GetUserByIDUseCase, deleteUser *userUseCase.DeleteUserUseCase, loginUser *userAuthUseCase.LoginUserUseCase, changePasswordUser *userAuthUseCase.ChangePasswordUseCase) *UserController {
	return &UserController{
		createUser:         createUser,
		getUserById:        getUserById,
		updateUser:         updateUser,
		deleteUser:         deleteUser,
		loginUser:          loginUser,
		changePasswordUser: changePasswordUser,
	}
}

// CreateUser godoc
// @Summary Criar novo usuário
// @Description Cria um novo usuário no sistema
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.CreateUserRequest true "Dados do usuário"
// @Success 201 {object} map[string]string "User created successfully"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users [post]
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userEntity, err := mapper.FromCreateRequestToUserEntity(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = ctrl.createUser.Execute(c.Request.Context(), userEntity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// GetUser godoc
// @Summary Buscar usuário por ID
// @Description Retorna os dados de um usuário específico
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ID do usuário"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 404 {object} map[string]string "User not found"
// @Router /users/{id} [get]
func (ctrl *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	userEntity, err := ctrl.getUserById.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	userResponse := mapper.FromUserEntityToUserResponse(userEntity)

	c.JSON(http.StatusOK, userResponse)
}

// UpdateUser godoc
// @Summary Atualizar usuário
// @Description Atualiza os dados de um usuário existente
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ID do usuário"
// @Param request body dto.UpdateUserRequest true "Dados atualizados"
// @Success 200 {object} map[string]string "User updated successfully"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users/{id} [put]
func (ctrl *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userEntity, err := mapper.FromUpdateRequestToUserEntity(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = ctrl.updateUser.Execute(c.Request.Context(), id, userEntity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser godoc
// @Summary Deletar usuário
// @Description Remove um usuário do sistema
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ID do usuário"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users/{id} [delete]
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	if err := ctrl.deleteUser.Execute(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// Login godoc
// @Summary Fazer login
// @Description Autentica o usuário e retorna tokens de acesso
// @Tags auth
// @Accept json
// @Produce json
// @Param request body securityDTO.LoginRequest true "Credenciais de login"
// @Success 200 {object} map[string]interface{} "user, access_token, refresh_token"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /auth/login [post]
func (ctrl *UserController) Login(c *gin.Context) {
	var req securityDTO.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginResponse, accessToken, refreshToken, err := ctrl.loginUser.Execute(c.Request.Context(), req.IDToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":          loginResponse,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// ChangePassword godoc
// @Summary Alterar senha
// @Description Altera a senha do usuário autenticado
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body securityDTO.ChangePasswordRequest true "Nova senha"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /auth/change-password [post]
func (ctrl *UserController) ChangePassword(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User UID not found in context"})
		return
	}

	userID, ok := uid.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UID in context is not a string"})
		return
	}

	var req securityDTO.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input := mapper.FromChangePasswordRequestToDomain(userID, &req)
	err := ctrl.changePasswordUser.Execute(c.Request.Context(), input.UserID, input.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
