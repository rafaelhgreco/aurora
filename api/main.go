package main

import (
	"log"

	// 1. Importa o nosso pacote builder
	"aurora.com/aurora-backend/internal/builder"
)

func main() {
	// 2. Chama a função Build para criar e injetar todas as dependências
	container, err := builder.Build()
	if err != nil {
		// Se o builder falhar (ex: não conseguir conectar ao DB no futuro),
		// a aplicação não deve iniciar.
		log.Fatalf("Falha ao inicializar as dependências da aplicação: %v", err)
	}

	// A função main não cria mais o roteador. Ela o recebe pronto do container.
	// O roteador já vem com todas as rotas (/v1/users, etc.) configuradas.
	
	log.Println("Servidor iniciando na porta :8080")

	// 3. Inicia o servidor usando o roteador que o container preparou
	if err := container.Router.Run(":8080"); err != nil {
		log.Fatalf("Falha ao iniciar o servidor Gin: %v", err)
	}
}