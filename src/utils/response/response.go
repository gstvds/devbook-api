package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(writer http.ResponseWriter, statusCode int, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(writer).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func Error(writer http.ResponseWriter, statusCode int, err error) {
	JSON(writer, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
