package repository

import (
	"archengine/internal/app/model"
	"archengine/internal/db"
)

func CreateUser(user *model.User) error {
	return db.DB.Create(user).Error
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := db.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
