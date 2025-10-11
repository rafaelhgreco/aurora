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
