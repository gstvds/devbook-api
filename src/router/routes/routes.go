package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI           string
	Method        string
	Exec          func(http.ResponseWriter, *http.Request)
	Authenticated bool
}

// Configure puts all routes insider the Router
func Configure(router *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoutes)

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Exec).Methods(route.Method)
	}

	return router
}
