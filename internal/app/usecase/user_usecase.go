package usecase

import (
	"archengine/dto"
	"archengine/internal/app/model"
	"archengine/internal/app/repository"
	"archengine/internal/util"
	"errors"
)

func RegisterUser(request dto.RegisterRequest) error {
	hashedPassword, _ := util.HashPassword(request.Password)
	user := model.User{Name: request.Name, Email: request.Email, Password: hashedPassword}
	return repository.CreateUser(&user)
}

func LoginUser(request dto.LoginRequest) (string, error) {
	user, err := repository.GetUserByEmail(request.Email)
	if err != nil || !util.CheckPasswordHash(request.Password, user.Password) {
		return "", errors.New("invalid credentials")
	}
	return util.GenerateJWT(user.Email)
}
