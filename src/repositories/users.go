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

// GetByEmail return a User from database by its email
func (repository Users) GetByEmail(email string) (models.User, error) {
	var user models.User

	err := repository.db.Select("id", "email", "password").Where("email = ?", email).First(&user).Error

	return user, err
}

// Follow allow a User to follow another
func (repository Users) Follow(userId, followerId uint64) error {
	var follower = models.Follower{UserId: userId, FollowerId: followerId}
	if err := repository.db.Create(&follower).Error; err != nil {
		return err
	}

	return nil
}

// Unfollow allow a User to unfollow another
func (repository Users) Unfollow(userId, followerId uint64) error {
	var databaseFollower models.Follower
	databaseFollower.FollowerId = followerId
	err := repository.db.Delete(&databaseFollower, "user_id = ?", userId).Error

	return err
}

// GetFollowers gets all followers form a User
func (repository Users) GetFollowers(userId uint64) ([]models.User, error) {
	var user []models.User

	err := repository.db.Joins("JOIN followers on users.id=followers.follower_id").Where("followers.user_id = ?", userId).Select("id", "username", "name", "email").Find(&user).Error

	fmt.Println(user)
	return user, err
}

// GetFollowing gets users a specific User is following
func (repository Users) GetFollowing(userId uint64) ([]models.User, error) {
	var user []models.User

	err := repository.db.Joins("JOIN followers on users.id=followers.user_id").Where("followers.follower_id = ?", userId).Select("id", "username", "name", "email").Find(&user).Error

	fmt.Println(user)
	return user, err
}
