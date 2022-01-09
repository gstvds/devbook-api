package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON sends a success response to the client
func JSON(writer http.ResponseWriter, statusCode int, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(writer).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

// Error sends a error response to the client
func Error(writer http.ResponseWriter, statusCode int, err error) {
	JSON(writer, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
