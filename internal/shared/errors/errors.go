package errors

import "errors"

// Domain errors
var (
    ErrNotFound          = errors.New("resource not found")
    ErrInvalidInput      = errors.New("invalid input")
    ErrUnauthorized      = errors.New("unauthorized")
    ErrForbidden         = errors.New("forbidden")
    ErrInternalServer    = errors.New("internal server error")
    ErrConflict          = errors.New("resource already exists")
    ErrBadRequest        = errors.New("bad request")
    ErrUnprocessable     = errors.New("unprocessable entity")
)

// User errors
var (
    ErrUserNotFound      = errors.New("user not found")
    ErrUserAlreadyExists = errors.New("user already exists")
    ErrInvalidCredentials = errors.New("invalid credentials")
)

// Event errors
var (
    ErrEventNotFound     = errors.New("event not found")
    ErrNoTicketsAvailable = errors.New("no tickets available")
)

// Order errors
var (
    ErrOrderNotFound     = errors.New("order not found")
    ErrOrderNotPending   = errors.New("order is not pending")
)

// Ticket errors
var (
    ErrTicketNotFound    = errors.New("ticket not found")
    ErrInvalidTicket     = errors.New("invalid ticket")
)