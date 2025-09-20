package controller

import (
	"net/http"

	"aurora.com/aurora-backend/internal/features/user/dto"
	"github.com/gin-gonic/gin"

	securityDTO "aurora.com/aurora-backend/internal/features/user/security/dto"
	userservice "aurora.com/aurora-backend/internal/features/user/service"
)

type UserController struct {
	userService userservice.IUserService
}

func NewUserController(service userservice.IUserService ) *UserController {
	return &UserController{
		userService: service,
	}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := ctrl.userService.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, userResponse)
}

func (ctrl *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	userResponse, err := ctrl.userService.FindByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

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

	userResponse, err := ctrl.userService.Update(c.Request.Context(), id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userResponse)
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	if err := ctrl.userService.Delete(c.Request.Context(), id); err != nil {
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

	loginResponse, err := ctrl.userService.Login(c.Request.Context(), req.IDToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loginResponse)
}

func (ctrl *UserController) ChangePassword(c *gin.Context) {
	// O UID não vem mais da URL. Ele é injetado de forma segura pelo middleware.
	uid, exists := c.Get("uid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User UID not found in context"})
		return
	}
	
	// Convertemos o uid (que é interface{}) para string
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

	if err := ctrl.userService.ChangePassword(c.Request.Context(), userID, &req); err != nil {
		// ...
		return
	}

	c.Status(http.StatusNoContent)
}