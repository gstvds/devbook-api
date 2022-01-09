package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		URI:           "/posts",
		Method:        http.MethodPost,
		Exec:          controllers.CreatePost,
		Authenticated: true,
	},
	{
		URI:           "/posts",
		Method:        http.MethodGet,
		Exec:          controllers.ListPost,
		Authenticated: true,
	},
	{
		URI:           "/posts/{postId}",
		Method:        http.MethodGet,
		Exec:          controllers.GetPost,
		Authenticated: true,
	},
	{
		URI:           "/posts/{postId}",
		Method:        http.MethodPut,
		Exec:          controllers.UpdatePost,
		Authenticated: true,
	},
	{
		URI:           "/posts/{postId}",
		Method:        http.MethodDelete,
		Exec:          controllers.DeletePost,
		Authenticated: true,
	},
}
