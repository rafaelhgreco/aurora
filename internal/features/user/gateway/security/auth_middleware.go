package security

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware cria um middleware do Gin para verificar o token do Firebase.
func AuthMiddleware(authClient *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Pega o token do cabeçalho "Authorization"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		// O formato esperado é "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}
		idToken := tokenParts[1]

		// 2. Verifica o token com o Firebase
		token, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// 3. Token é válido! Adiciona o UID ao contexto do Gin para uso posterior.
		c.Set("uid", token.UID)

		// 4. Continua para o próximo handler na cadeia.
		c.Next()
	}
}