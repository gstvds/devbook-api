package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/providers/hash_provider"
	"api/src/repositories"
	"api/src/utils/response"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.Error(writer, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()

	repository := repositories.NewUserRepository(db)

	databaseUser, err := repository.GetByEmail(user.Email)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	if err = hash_provider.CheckPassword(user.Password, databaseUser.Password); err != nil {
		customErr := errors.New("invalid username or password")
		response.Error(writer, http.StatusUnauthorized, customErr)
		return
	}
}
