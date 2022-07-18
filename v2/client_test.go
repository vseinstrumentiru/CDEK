package v2

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func createTestClient() Client {
	wd, _ := os.Getwd()
	godotenv.Load(fmt.Sprintf("%s/.env.local", wd))

	clientId := os.Getenv("CDEK_CLIENT_ID")
	clientSecretId := os.Getenv("CDEK_SECRET_ID")

	// public cdek test credentials
	if clientId == "" {
		clientId = "EMscd6r9JnFiQ3bLoyjJY6eM78JrJceI"
	}
	if clientSecretId == "" {
		clientSecretId = "PjLZkKBHEiLK3YsjtNrt3TGNG0ahs3kG"
	}

	return NewClient(&Options{
		Endpoint: EndpointTest,
		Credentials: &Credentials{
			ClientID:     os.Getenv("CDEK_CLIENT_ID"),
			ClientSecret: os.Getenv("CDEK_SECRET_ID"),
		},
	})
}
