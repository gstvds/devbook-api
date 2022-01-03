package repositories

import (
	"api/src/models"
	"errors"
	"time"

	adminFirestore "cloud.google.com/go/firestore"
	adminAuth "firebase.google.com/go/auth"
	"google.golang.org/api/iterator"
)

type UserRepository interface {
	CreateUser(user models.User) error
	GetUserById(userId string) (models.User, error)
	GetAllByUsername(username string) ([]models.User, error)
	UpdateUser(userId string, user models.User) error
}

type repository struct{}

func NewUserRepository() UserRepository {
	return &repository{}
}

func (repository *repository) CreateUser(user models.User) error {
	app := NewFirebaseApp()
	firestore, err := app.GetFirestore()
	if err != nil {
		return err
	}

	defer firestore.Close()

	auth, err := app.GetAuth()
	if err != nil {
		return err
	}

	getUser, _ := auth.GetUserByEmail(Ctx, user.Email)
	if getUser != nil {
		return errors.New("user already exists")
	}

	iter := firestore.Collection("users").Where("username", "==", user.Password).Documents(Ctx)
	for {
		document, err := iter.Next()
		if err == iterator.Done || !document.Exists() {
			break
		} else {
			return errors.New("user already exists")
		}
	}

	params := (&adminAuth.UserToCreate{}).Email(user.Email).Password(user.Password).DisplayName(user.Name).Disabled(false)
	createUser, err := auth.CreateUser(Ctx, params)
	if err != nil {
		return err
	}

	var search_username []string
	temp := ""
	for _, data := range user.Username {
		temp = temp + string(data)
		search_username = append(search_username, string(temp))
	}

	_, err = firestore.Collection("users").Doc(createUser.UID).Create(Ctx, map[string]interface{}{
		"name":            user.Name,
		"email":           user.Email,
		"username":        user.Username,
		"search_username": search_username,
		"created_at":      time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (repository *repository) GetUserById(userId string) (models.User, error) {
	app := NewFirebaseApp()
	firestore, err := app.GetFirestore()
	if err != nil {
		return models.User{}, err
	}

	defer firestore.Close()

	var user models.User

	dbUser, err := firestore.Collection("users").Doc(userId).Get(Ctx)
	if err != nil {
		return models.User{}, err
	}

	dbUser.DataTo(&user)
	user.Id = userId

	return user, nil
}

func (repository *repository) GetAllByUsername(username string) ([]models.User, error) {
	app := NewFirebaseApp()
	firestore, err := app.GetFirestore()
	if err != nil {
		return nil, err
	}

	defer firestore.Close()

	var users []models.User

	iter := firestore.Collection("users").Where("search_username", "array-contains", username).Documents(Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return nil, err
		}

		var parseUser models.User
		doc.DataTo(&parseUser)

		users = append(users, parseUser)
	}

	return users, nil
}

func (repository *repository) UpdateUser(userId string, user models.User) error {
	app := NewFirebaseApp()
	firestore, err := app.GetFirestore()
	if err != nil {
		return err
	}

	defer firestore.Close()

	var updateData []adminFirestore.Update

	if string(user.Username) != "" {
		var updateUsername = adminFirestore.Update{
			Path:  "username",
			Value: user.Username,
		}
		updateData = append(updateData, updateUsername)

		var search_username []string
		temp := ""
		for _, data := range user.Username {
			temp = temp + string(data)
			search_username = append(search_username, string(temp))
		}

		var updateSearchUsername = adminFirestore.Update{
			Path:  "search_username",
			Value: search_username,
		}
		updateData = append(updateData, updateSearchUsername)
	}

	if string(user.Name) != "" {
		var updateName = adminFirestore.Update{
			Path:  "name",
			Value: user.Name,
		}
		updateData = append(updateData, updateName)
	}

	var updatedAt = adminFirestore.Update{
		Path:  "updated_at",
		Value: time.Now(),
	}

	updateData = append(updateData, updatedAt)

	_, err = firestore.Collection("users").Doc(userId).Update(Ctx, updateData)
	if err != nil {
		return err
	}

	return nil
}
