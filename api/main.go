package main

import (
	"log"

	"aurora.com/aurora-backend/internal/builder"
)

func main() {
	container, err := builder.Build()
	if err != nil {
		log.Fatalf("Falha ao inicializar as dependências da aplicação: %v", err)
	}
	
	log.Println("Servidor iniciando na porta :8080")

	if err := container.Router.Run(":8080"); err != nil {
		log.Fatalf("Falha ao iniciar o servidor Gin: %v", err)
	}
}