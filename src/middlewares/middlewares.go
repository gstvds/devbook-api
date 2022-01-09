package middlewares

import (
	"api/src/middlewares/authentication"
	"api/src/utils/response"
	"log"
	"net/http"
)

// Logger adds a initial log with the received method, URI and host
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("\n [%s] %s %s", request.Method, request.RequestURI, request.Host)

		next(writer, request)
	}
}

// Authenticate a User
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := authentication.ValidateToken(request); err != nil {
			response.Error(writer, http.StatusUnauthorized, err)
			return
		}

		next(writer, request)
	}
}
