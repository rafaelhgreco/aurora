package controller

import (
	"net/http"

	"aurora.com/aurora-backend/internal/features/order/dto"
	"aurora.com/aurora-backend/internal/features/order/mapper"
	usecase "aurora.com/aurora-backend/internal/features/order/use-case"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderUseCase *usecase.CreateOrderUseCase
}

func NewOrderController(orderUseCase *usecase.CreateOrderUseCase) *OrderController {
	return &OrderController{orderUseCase: orderUseCase}
}

func (ctrl *OrderController) CreateOrder(c *gin.Context) {
	var req dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderDomain, err := mapper.FromCreateOrderRequestToDomain(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	orderId, err := ctrl.orderUseCase.Execute(c.Request.Context(), orderDomain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order_id": orderId})
}
