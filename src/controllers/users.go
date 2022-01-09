package controllers

import (
	"api/src/database"
	"api/src/middlewares/authentication"
	"api/src/models"
	"api/src/repositories"
	"api/src/utils/response"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Create a new User
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

// List all Users by username
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

// Get a User by userId
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

// Update a User by userId
func Update(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	tokenUserId, err := authentication.ExtractUserId(request)
	if err != nil {
		response.Error(writer, http.StatusUnauthorized, err)
		return
	}

	if userId != tokenUserId {
		response.Error(writer, http.StatusForbidden, errors.New("you can not update a user that is not yours"))
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

// Delete a User by userId
func Delete(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	tokenUserId, err := authentication.ExtractUserId(request)
	if err != nil {
		response.Error(writer, http.StatusUnauthorized, err)
		return
	}

	if userId != tokenUserId {
		response.Error(writer, http.StatusForbidden, errors.New("you can not delete a user that is not yours"))
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

// Follow allow a User to follow another
func Follow(writer http.ResponseWriter, request *http.Request) {
	followerId, err := authentication.ExtractUserId(request)
	if err != nil {
		response.Error(writer, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(request)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	if followerId == userId {
		response.Error(writer, http.StatusForbidden, errors.New("invalid operation. you can't follow yourself"))
		return
	}

	db := database.GetDB()
	repository := repositories.NewUserRepository(db)

	if err = repository.Follow(userId, followerId); err != nil {
		response.Error(writer, http.StatusInternalServerError, err)
		return
	}

	response.JSON(writer, http.StatusNoContent, nil)
}

// Unfollow allow a User to unfollow another
func Unfollow(writer http.ResponseWriter, request *http.Request) {
	followerId, err := authentication.ExtractUserId(request)
	if err != nil {
		response.Error(writer, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(request)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	if followerId == userId {
		response.Error(writer, http.StatusForbidden, errors.New("invalid operation. you can't unfollow yourself"))
	}

	db := database.GetDB()
	repository := repositories.NewUserRepository(db)
	if err = repository.Unfollow(userId, followerId); err != nil {
		response.Error(writer, http.StatusInternalServerError, err)
		return
	}

	response.JSON(writer, http.StatusNoContent, nil)
}

// Followers gets all followers form a User
func Followers(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()
	repository := repositories.NewUserRepository(db)
	followers, err := repository.GetFollowers(userId)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err)
		return
	}

	response.JSON(writer, http.StatusOK, followers)
}

// Following gets users a specific User is following
func Following(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()
	repository := repositories.NewUserRepository(db)
	followers, err := repository.GetFollowing(userId)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err)
		return
	}

	response.JSON(writer, http.StatusOK, followers)
}
