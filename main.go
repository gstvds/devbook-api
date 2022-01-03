package main

import (
	"api/src/repositories"
	"api/src/router"
	"api/src/utils/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()

	router := router.Create()

	fmt.Printf("Server running on PORT %d\n", config.PORT)
	_, err := repositories.GetApp()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), router))
}
