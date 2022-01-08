package main

import (
	"api/src/database"
	"api/src/router"
	"api/src/utils/config"
	"fmt"
	"log"
	"net/http"
)

func listen() {
	routes := router.Create()

	fmt.Printf("Server running on PORT %d\n", config.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), routes))
}

func main() {
	config.LoadEnv()
	database.Setup()

	listen()
}
