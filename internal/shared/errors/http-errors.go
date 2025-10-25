package errors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
    Error   string `json:"error"`
    Message string `json:"message,omitempty"`
}

// HandleError mapeia erros de domínio para respostas HTTP
func HandleError(c *gin.Context, err error) {
    var statusCode int
    var message string

    switch {
    case errors.Is(err, ErrNotFound), 
         errors.Is(err, ErrUserNotFound), 
         errors.Is(err, ErrEventNotFound),
         errors.Is(err, ErrOrderNotFound),
         errors.Is(err, ErrTicketNotFound):
        statusCode = http.StatusNotFound
        message = err.Error()

    case errors.Is(err, ErrUnauthorized), 
         errors.Is(err, ErrInvalidCredentials):
        statusCode = http.StatusUnauthorized
        message = err.Error()

    case errors.Is(err, ErrForbidden):
        statusCode = http.StatusForbidden
        message = err.Error()

    case errors.Is(err, ErrInvalidInput), 
         errors.Is(err, ErrBadRequest):
        statusCode = http.StatusBadRequest
        message = err.Error()

    case errors.Is(err, ErrConflict), 
         errors.Is(err, ErrUserAlreadyExists):
        statusCode = http.StatusConflict
        message = err.Error()

    case errors.Is(err, ErrUnprocessable),
         errors.Is(err, ErrNoTicketsAvailable),
         errors.Is(err, ErrOrderNotPending),
         errors.Is(err, ErrInvalidTicket):
        statusCode = http.StatusUnprocessableEntity
        message = err.Error()

    default:
        statusCode = http.StatusInternalServerError
        message = "internal server error"
    }

    c.JSON(statusCode, ErrorResponse{
        Error:   http.StatusText(statusCode),
        Message: message,
    })
}