package firebase

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

// FirebaseApp encapsula o cliente do Firebase para ser injetado.
type FirebaseApp struct {
	*firebase.App
}

// NewFirebaseApp inicializa e retorna uma nova conexão com o Firebase.
func NewFirebaseApp(ctx context.Context) (*FirebaseApp, error) {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Printf("Aviso: Não foi possível carregar o arquivo .env. Usando variáveis de ambiente do sistema, se disponíveis.")
	}

	credentialsPath := os.Getenv("FIREBASE_CREDENTIALS_FILE")
	if credentialsPath == "" {
		log.Fatal("A variável de ambiente FIREBASE_CREDENTIALS_FILE não está definida.")
	}

	opt := option.WithCredentialsFile(credentialsPath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	return &FirebaseApp{app}, nil
}
