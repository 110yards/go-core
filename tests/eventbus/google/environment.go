package google

import (
	"os"

	"110yards.ca/libs/go/core/logger"
	"github.com/joho/godotenv"
)

var projectId string

func setup() {
	err := godotenv.Load("../../../.env")

	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	projectId = os.Getenv("GCLOUD_TEST_PROJECT")
}
