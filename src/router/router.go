package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

func Create() *mux.Router {
	router := mux.NewRouter()
	return routes.Configure(router)
}
