package users

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID          uint   `gorm:"primary_key"`
	Username    string `gorm:"type:varchar(50)"`
	Email       string `gorm:"type:varchar(255);unique_index"`
	PhoneNumber string `gorm:"type:varchar(11);unique_index"`
	UpdatedAt   time.Time
	CreatedAt   time.Time
}