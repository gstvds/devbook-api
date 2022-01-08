package repositories

import (
	"api/src/models"
	"fmt"
	"gorm.io/gorm"
)

type Users struct {
	db *gorm.DB
}

// NewUserRepository returns a User repository to access the database
func NewUserRepository(db *gorm.DB) *Users {
	return &Users{db}
}

// Create a User in the database
func (repository Users) Create(user models.User) error {
	if err := repository.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// Get a specific User by its userId
func (repository Users) Get(userId uint64) (models.User, error) {
	var user models.User

	err := repository.db.Select("id", "name", "username", "email").Where("id = ?", userId).First(&user).Error

	return user, err
}

// List all users by username
func (repository Users) List(username string) ([]models.User, error) {
	var users []models.User

	// Format the username to be able to find any username that looks like
	username = fmt.Sprintf("%%%s%%", username)

	err := repository.db.Select("id", "name", "username", "email").Where("username ILIKE ?", username).Find(&users).Error

	return users, err
}

// Update a specific User using its userId
func (repository Users) Update(userId uint64, user models.User) error {
	var databaseUser models.User
	databaseUser.Id = userId

	repository.db.First(&databaseUser)

	databaseUser.Name = user.Name
	databaseUser.Email = user.Email
	databaseUser.Username = user.Username

	err := repository.db.Save(&databaseUser).Error

	return err
}

// Delete a User from database using its userId
func (repository Users) Delete(userId uint64) error {
	var databaseUser models.User
	databaseUser.Id = userId

	err := repository.db.Delete(&databaseUser).Error

	return err
}
