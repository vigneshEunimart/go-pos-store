package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// to get the env variables
func GetEnvVAlues(s string) string {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("error while getting the values from env")
		log.Fatal(err)
	}
	return os.Getenv(s)
}
