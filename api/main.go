package main

import (
	"log"

	"aurora.com/aurora-backend/internal/builder"
	"aurora.com/aurora-backend/internal/shared/logger"
)

func main() {
	logger.Init()
	container, err := builder.Build()
	if err != nil {
		log.Fatalf("Falha ao inicializar as dependências da aplicação: %v", err)
	}

	logger.Info("Servidor iniciando na porta :8080")
	logger.Info("Swagger disponível em http://localhost:8080/swagger/index.html")

	if err := container.Router.Run(":8080"); err != nil {
		log.Fatalf("Falha ao iniciar o servidor Gin: %v", err)
	}
}
