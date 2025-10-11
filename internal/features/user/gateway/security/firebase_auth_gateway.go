package security

import (
	"context"
	"errors"
	"os"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/golang-jwt/jwt/v5"

	"aurora.com/aurora-backend/internal/features/user/domain"
)

// Your JWT secret key. This MUST be stored securely, for example, in a .env file.
var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// FirebaseAuthGateway is our adapter. It wraps the real Firebase client.
type FirebaseAuthGateway struct {
	firebaseAuth *auth.Client
}

// NewFirebaseAuthGateway is the constructor for our adapter.
func NewFirebaseAuthGateway(client *auth.Client) domain.AuthClient {
	return &FirebaseAuthGateway{firebaseAuth: client}
}

// VerifyIDToken simply delegates the call to the real Firebase client.
func (g *FirebaseAuthGateway) VerifyIDToken(ctx context.Context, idToken string) (string, error) {
	token, err := g.firebaseAuth.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", err
	}
	return token.UID, nil
}

// GenerateAccessToken is our custom implementation.
func (g *FirebaseAuthGateway) GenerateAccessToken(ctx context.Context, userID string) (string, error) {
	// Create the token claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(), // Expires in 1 hour
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GenerateRefreshToken is our custom implementation.
func (g *FirebaseAuthGateway) GenerateRefreshToken(ctx context.Context, userID string) (string, error) {
	// Refresh tokens typically have a much longer expiration time
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(), // Expires in 30 days
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (g *FirebaseAuthGateway) CreateUser(ctx context.Context, user *domain.User) (string, error) {
	params := (&auth.UserToCreate{}).
		Email(user.Email).
		Password(user.Password).
		DisplayName(user.Name)

	// Chama a função real do SDK do Firebase com os parâmetros convertidos.
	userRecord, err := g.firebaseAuth.CreateUser(ctx, params)
	if err != nil {
		return "", err
	}

	// Retorna o UID.
	return userRecord.UID, nil
}

func (g *FirebaseAuthGateway) UpdateUser(ctx context.Context, uid string, params interface{}) (interface{}, error) {
	// We need to cast the params to the type the Firebase SDK expects.
	// This makes our gateway flexible.
	updateParams, ok := params.(*auth.UserToUpdate)
	if !ok {
		return nil, errors.New("invalid params type for UpdateUser; expected *auth.UserToUpdate")
	}

	userRecord, err := g.firebaseAuth.UpdateUser(ctx, uid, updateParams)
	if err != nil {
		return nil, err
	}

	return userRecord, nil
}
