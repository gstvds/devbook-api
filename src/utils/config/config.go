package config

import (
	"api/src/utils"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type database struct {
	User     string `json:"db_user,omitempty"`
	Password string `json:"db_password,omitempty"`
	Host     string `json:"db_host,omitempty"`
	Name     string `json:"db_name,omitempty"`
	Port     uint   `json:"db_port,omitempty"`
}

var (
	DATABASE   database
	SECRET_KEY []byte
	PORT       = 0
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

	SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

	utils.ParseJSONFile("./.credentials/database.json", &DATABASE)
}
