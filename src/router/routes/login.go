package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginRoutes = Route{
	URI:           "/login",
	Method:        http.MethodPost,
	Exec:          controllers.Login,
	Authenticated: false,
}
