package controllers

import (
	"api/src/database"
	"api/src/middlewares/authentication"
	"api/src/models"
	"api/src/providers/hash_provider"
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

// CreateUser a new User
func CreateUser(writer http.ResponseWriter, request *http.Request) {
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

// ListUser all Users by username
func ListUser(writer http.ResponseWriter, request *http.Request) {
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

// GetUser a User by userId
func GetUser(writer http.ResponseWriter, request *http.Request) {
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

// UpdateUser a User by userId
func UpdateUser(writer http.ResponseWriter, request *http.Request) {
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

// DeleteUser a User by userId
func DeleteUser(writer http.ResponseWriter, request *http.Request) {
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

// FollowUser allow a User to follow another
func FollowUser(writer http.ResponseWriter, request *http.Request) {
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

// UnfollowUser allow a User to unfollow another
func UnfollowUser(writer http.ResponseWriter, request *http.Request) {
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

// GetFollowers gets all followers form a User
func GetFollowers(writer http.ResponseWriter, request *http.Request) {
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

// GetFollowing gets users a specific User is following
func GetFollowing(writer http.ResponseWriter, request *http.Request) {
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

// UpdatePassword updates a User password
func UpdatePassword(writer http.ResponseWriter, request *http.Request) {
	userIdFromToken, err := authentication.ExtractUserId(request)
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

	if userIdFromToken != userId {
		response.Error(writer, http.StatusForbidden, errors.New("it's not possible to update a user that is not yourself"))
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.Error(writer, http.StatusUnprocessableEntity, err)
	}

	var password models.Password

	if err = json.Unmarshal(body, &password); err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()
	repository := repositories.NewUserRepository(db)
	databasePassword, err := repository.GetPassword(userId)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err)
		return
	}

	if err = hash_provider.CheckPassword(password.Old, databasePassword); err != nil {
		response.Error(writer, http.StatusUnauthorized, errors.New("invalid password"))
	}

	hashedPassword, err := hash_provider.Hash(password.New)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err)
		return
	}

	if err = repository.UpdatePassword(userId, string(hashedPassword)); err != nil {
		response.Error(writer, http.StatusInternalServerError, err)
		return
	}

	response.JSON(writer, http.StatusNoContent, nil)
}
