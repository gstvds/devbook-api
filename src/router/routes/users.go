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
		Exec:          controllers.ListUser,
		Authenticated: true,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodGet,
		Exec:          controllers.GetUser,
		Authenticated: true,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodPut,
		Exec:          controllers.UpdateUser,
		Authenticated: true,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodDelete,
		Exec:          controllers.DeleteUser,
		Authenticated: false,
	},
	{
		URI:           "/users/{userId}/follow",
		Method:        http.MethodPost,
		Exec:          controllers.FollowUser,
		Authenticated: true,
	},
	{
		URI:           "/users/{userId}/unfollow",
		Method:        http.MethodPost,
		Exec:          controllers.UnfollowUser,
		Authenticated: true,
	},
	{
		URI:           "/users/{userId}/followers",
		Method:        http.MethodGet,
		Exec:          controllers.GetFollowers,
		Authenticated: true,
	},
	{
		URI:           "/users/{userId}/following",
		Method:        http.MethodGet,
		Exec:          controllers.GetFollowing,
		Authenticated: true,
	},
	{
		URI:           "/users/{userId}/update_password",
		Method:        http.MethodPost,
		Exec:          controllers.UpdatePassword,
		Authenticated: true,
	},
}
