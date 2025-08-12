package controller

import (
	"net/http"
	"strconv"

	"aurora.com/aurora-backend/internal/features/user/dto"
	"github.com/gin-gonic/gin"

	userservice "aurora.com/aurora-backend/internal/features/user/service"
)

type UserController struct {
	userService userservice.IUserService
}

func NewUserController(service userservice.IUserService) *UserController {
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

	// O controller agora só chama o método do serviço.
	// Toda a lógica de mapeamento e orquestração foi movida.
	userResponse, err := ctrl.userService.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, userResponse)
}

func (ctrl *UserController) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Delega diretamente para o serviço.
	userResponse, err := ctrl.userService.FindByID(c.Request.Context(), id)
	if err != nil {
		// O serviço já lidou com a lógica, aqui só tratamos o resultado.
		// Poderíamos ter erros mais específicos para retornar 404 vs 500.
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userResponse)
}