package cdek

import (
	"os"
)

const JsonContentType = "application/json"

func GetServerUrl() string {
	return os.Getenv("SERVER_URL")
}
