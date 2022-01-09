package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:           "/users",
		Method:        http.MethodPost,
		Exec:          controllers.Create,
		Authenticated: false,
	},
	{
		URI:           "/users",
		Method:        http.MethodGet,
		Exec:          controllers.List,
		Authenticated: true,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodGet,
		Exec:          controllers.Get,
		Authenticated: true,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodPut,
		Exec:          controllers.Update,
		Authenticated: true,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodDelete,
		Exec:          controllers.Delete,
		Authenticated: false,
	},
}
