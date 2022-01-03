package controllers

import (
	"api/src/models"
	"api/src/repositories"
	"api/src/utils/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.Error(writer, http.StatusUnprocessableEntity, err)
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		response.Error(writer, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewUserRepository()

	err = repository.CreateUser(user)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	response.JSON(writer, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "user created successfully",
	})
}

func GetAllUsers(writer http.ResponseWriter, request *http.Request) {
	username := strings.ToLower(request.URL.Query().Get("user"))
	repository := repositories.NewUserRepository()

	users, err := repository.GetAllByUsername(username)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	response.JSON(writer, http.StatusOK, users)
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userId := params["userId"]
	repository := repositories.NewUserRepository()

	user, err := repository.GetUserById(userId)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	response.JSON(writer, http.StatusOK, user)
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userId := params["userId"]

	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.Error(writer, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	repository := repositories.NewUserRepository()
	err = repository.UpdateUser(userId, user)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	response.JSON(writer, http.StatusNoContent, nil)
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Deleting a User"))
}
