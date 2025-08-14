package repository

import (
	"context"
	"fmt"
	"sync"
	"time"

	"aurora.com/aurora-backend/internal/features/user/domain"
)

type UserInMemoryRepository struct {
	mu      sync.Mutex
	users   map[int]*domain.User
	counter int
}

func NewUserInMemoryRepository() domain.UserRepository {
	return &UserInMemoryRepository{
		users:   make(map[int]*domain.User),
		counter: 0,
	}
}

func (r *UserInMemoryRepository) Save(ctx context.Context, user *domain.User) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.counter++
	user.ID = r.counter
	user.CreatedAt = time.Now()
	r.users[user.ID] = user

	return user, nil
}

func (r *UserInMemoryRepository) FindByID(ctx context.Context, id int) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user with id %d not found", id)
	}

	return user, nil
}