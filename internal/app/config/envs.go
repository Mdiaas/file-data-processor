package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	envGcsCredentials  = "GOOGLE_CREDENTIALS"
	envBucketName      = "BUCKET_NAME"
	envFileUploadTopic = "FILE_UPLOAD_TOPIC"
	envFileUploadSub   = "FILE_UPLOAD_SUBSCRIBER"
	envNumberOfWorkers = "NUMBER_OF_WORKS"
	envAPIServerHost   = "API_SERVER_HOST"
	envProjectId       = "PROJECT_ID"
)

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}
	envWorkers := os.Getenv(envNumberOfWorkers)
	workers, err := strconv.Atoi(envWorkers)
	if err != nil {
		log.Fatal("failed to get number of workers")
	}

	return Config{
		Google: Google{
			Credentials: os.Getenv(envGcsCredentials),
		},
		CloudStorage: CloudStorage{
			BucketName: os.Getenv(envBucketName),
		},
		Topics: Topics{
			FileUpload: os.Getenv(envFileUploadTopic),
		},
		Subscribers: Subscribers{
			FileUpload: os.Getenv(envFileUploadSub),
		},
		Workers: Workers{
			NumberOfWorks: workers,
		},
		API: API{
			ServerHost: os.Getenv(envAPIServerHost),
		},
		Project: Project{
			Id: os.Getenv(envProjectId),
		},
	}
}
