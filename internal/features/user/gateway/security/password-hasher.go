package security

import (
	"context"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

const defaultBcryptCost = 12

type BcryptHasher struct {
    Cost int
}

// NewBcryptHasher usa (na ordem):
// 1. Valor passado explicitamente (>0)
// 2. Variável de ambiente AURORA_BCRYPT_COST (se válida)
// 3. defaultBcryptCost
func NewBcryptHasher(overrides ...int) *BcryptHasher {
    cost := 0
    if len(overrides) > 0 && overrides[0] > 0 {
        cost = overrides[0]
    }
    if cost == 0 {
        if envVal := os.Getenv("AURORA_BCRYPT_COST"); envVal != "" {
            if parsed, err := strconv.Atoi(envVal); err == nil && parsed > 0 {
                cost = parsed
            }
        }
    }
    if cost == 0 {
        cost = defaultBcryptCost
    }
    return &BcryptHasher{Cost: cost}
}

func (h *BcryptHasher) Hash(ctx context.Context, plain string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(plain), h.Cost)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}