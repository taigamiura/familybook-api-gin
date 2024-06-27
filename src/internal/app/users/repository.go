package users

import (
	"errors"

	"github.com/familybook-project/familybook-api-gin/src/internal/db"
	"gorm.io/gorm"
)

func GetAllUsers() ([]User, error) {
	var users []User
	if err := db.DBInstance.DB.Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users, errors.New("Users not found")
		}
		return users, err
	}
	return users, nil
}

func GetUserById(id string) (User, error) {
	var user User
	if err := db.DBInstance.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("User not found")
		}
		return user, err
	}
	return user, nil
}
