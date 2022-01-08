package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/utils/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func Create(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.Error(writer, http.StatusUnprocessableEntity, err)
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		response.Error(writer, http.StatusInternalServerError, err)
		return
	}

	if err = user.Validate("register"); err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()
	repository := repositories.NewUserRepository(db)

	if err = repository.Create(user); err != nil {
		response.Error(writer, http.StatusInternalServerError, err)
		return
	}

	response.JSON(writer, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "user created successfully",
	})
}

func List(writer http.ResponseWriter, request *http.Request) {
	username := strings.ToLower(request.URL.Query().Get("user"))

	db := database.GetDB()
	repository := repositories.NewUserRepository(db)

	users, err := repository.List(username)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	response.JSON(writer, http.StatusOK, users)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()
	repository := repositories.NewUserRepository(db)

	user, err := repository.Get(userId)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err)
		return
	}

	response.JSON(writer, http.StatusOK, user)
}

func Update(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

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

	if err = user.Validate("update"); err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()
	repository := repositories.NewUserRepository(db)

	if err = repository.Update(userId, user); err != nil {
		response.Error(writer, http.StatusInternalServerError, err)
		return
	}

	response.JSON(writer, http.StatusNoContent, nil)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()
	repository := repositories.NewUserRepository(db)

	if err = repository.Delete(userId); err != nil {
		response.Error(writer, http.StatusInternalServerError, err)
		return
	}

	response.JSON(writer, http.StatusNoContent, nil)
}
