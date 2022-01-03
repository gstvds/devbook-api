package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:           "/users",
		Method:        http.MethodPost,
		Exec:          controllers.CreateUser,
		Authenticated: false,
	},
	{
		URI:           "/users",
		Method:        http.MethodGet,
		Exec:          controllers.GetAllUsers,
		Authenticated: false,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodGet,
		Exec:          controllers.GetUser,
		Authenticated: false,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodPut,
		Exec:          controllers.UpdateUser,
		Authenticated: false,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodDelete,
		Exec:          controllers.DeleteUser,
		Authenticated: false,
	},
}
