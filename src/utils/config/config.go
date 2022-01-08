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
	Port     uint `json:"db_port,omitempty"`
}

var (
	DATABASE database
	PORT     = 0
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

	utils.ParseJSONFile("./.credentials/database.json", &DATABASE)
}
