package services

import (
	"auth-api/models"
	"auth-api/repository"
	"errors"
)

func Signup(user models.User) error {

	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	return repository.CreateUser(user)
}

func Login(email, password string) error {

	return repository.Login(email, password)
}

func ChangePassword(email, oldPassword, newPassword string) error {

	return repository.ChangePassword(email, oldPassword, newPassword)
}

func Profile(email string) (*models.User, error) {

	return repository.GetProfile(email)
}

func Logout() error {

	return nil
}