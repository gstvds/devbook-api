package routes

import (
	"api/src/middlewares"
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
		if route.Authenticated {
			router.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Exec))).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI, middlewares.Logger(route.Exec)).Methods(route.Method)
		}
	}

	return router
}
