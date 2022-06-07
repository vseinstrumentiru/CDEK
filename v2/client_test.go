package v2

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func createTestClient() Client {
	wd, _ := os.Getwd()
	godotenv.Load(fmt.Sprintf("%s/.env.local", wd))

	return NewClient(&Options{
		Endpoint: EndpointTest,
		Credentials: &Credentials{
			ClientID:     os.Getenv("CDEK_CLIENT_ID"),
			ClientSecret: os.Getenv("CDEK_SECRET_ID"),
		},
	})
}
