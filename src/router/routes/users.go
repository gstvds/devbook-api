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
	{
		URI:           "/users/{userId}/follow",
		Method:        http.MethodPost,
		Exec:          controllers.Follow,
		Authenticated: true,
	},
	{
		URI:           "/users/{userId}/unfollow",
		Method:        http.MethodPost,
		Exec:          controllers.Unfollow,
		Authenticated: true,
	},
	{
		URI:           "/users/{userId}/followers",
		Method:        http.MethodGet,
		Exec:          controllers.Followers,
		Authenticated: true,
	},
	{
		URI:           "/users/{userId}/following",
		Method:        http.MethodGet,
		Exec:          controllers.Following,
		Authenticated: true,
	},
}
