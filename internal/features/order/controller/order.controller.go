package controller

import (
	"net/http"

	"aurora.com/aurora-backend/internal/features/order/dto"
	"aurora.com/aurora-backend/internal/features/order/mapper"
	usecase "aurora.com/aurora-backend/internal/features/order/use-case"
	sharedErrors "aurora.com/aurora-backend/internal/shared/errors"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderUseCase *usecase.CreateOrderUseCase
}

func NewOrderController(orderUseCase *usecase.CreateOrderUseCase) *OrderController {
	return &OrderController{orderUseCase: orderUseCase}
}
// CreateOrder godoc
// @Summary Criar nova ordem
// @Description Cria uma nova ordem de compra para um evento
// @Tags orders
// @Accept json
// @Produce json
// @Param request body dto.CreateOrderRequest true "Dados da ordem"
// @Success 201 {object} map[string]interface{} "message, order_id"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 422 {object} map[string]string "Unprocessable Entity"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /orders [post]
func (ctrl *OrderController) CreateOrder(c *gin.Context) {
	var req dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sharedErrors.HandleError(c, sharedErrors.ErrInvalidInput)
		return
	}

	orderDomain, err := mapper.FromCreateOrderRequestToDomain(&req)
	if err != nil {
		sharedErrors.HandleError(c, err)
		return
	}

	orderId, err := ctrl.orderUseCase.Execute(c.Request.Context(), orderDomain)
	if err != nil {
		sharedErrors.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order_id": orderId})
}
