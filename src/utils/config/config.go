package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DATABASE_URL    = ""
	PORT            = 0
	PROJECT_ID      = ""
	SERVICE_ACCOUNT []byte
)

// LoadEnv loads the environment variables
func LoadEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		PORT = 3000
	}

	PROJECT_ID = os.Getenv("PROJECT_ID")

	SERVICE_ACCOUNT, err = ioutil.ReadFile("./.credentials/serviceAccount.json")
	if err != nil {
		log.Fatal("Failed to parse config file")
	}

	DATABASE_URL = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
