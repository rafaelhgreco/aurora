package domain

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string // Adicionar hash
	CreatedAt time.Time
}