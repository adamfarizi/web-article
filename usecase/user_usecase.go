package usecase

import (
	"fmt"
	"strings"
	"web-article/model"
	"web-article/repository"
	"web-article/utils/service"
)

type userUseCase struct {
	repo repository.UserRepository
}

type UserUseCase interface {
	GetUserByEmail(email string) (model.UserLogin, error)
	CreateUser(user model.User) (model.User, error)
	GetAllUser(name string) ([]model.User, error)
	GetUserById(id int) (model.User, error)
	UpdateUser(user model.User, id int) (model.User, error)
	DeleteUser(userID int) error
	IsUserIdExists(id int) (bool, error)
	IsUserEmailExists(email string) (bool, error)
}

func (u *userUseCase) GetUserByEmail(email string) (model.UserLogin, error) {
	if email == "" {
		return model.UserLogin{}, fmt.Errorf("email are required")
	}

	return u.repo.GetUserByEmail(email)
}

func (u *userUseCase) CreateUser(user model.User) (model.User, error) {
	exists, err := u.repo.IsUserEmailExists(user.Email)
	if err != nil {
		return model.User{}, err
	}
	if exists {
		return model.User{}, fmt.Errorf("user with Email %s already exists", user.Email)
	}

	if user.Role == "" {
		return model.User{}, fmt.Errorf("role are required")
	}

	if user.Role != "admin" && user.Role != "user" && user.Role != "editor" {
		return model.User{}, fmt.Errorf("choose role between admin, user, or editor")
	}

	if len(user.Password) < 8 {
		return model.User{}, fmt.Errorf("password must be at least 8 characters")
	}

	hashedPassword, err := service.HashPassword(user.Password)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to hash password: %v", err)
	}
	user.Password = hashedPassword

	return u.repo.CreateUser(user)
}

func (u *userUseCase) GetAllUser(name string) ([]model.User, error) {
	nameLower := strings.ToLower(name)
	return u.repo.GetAllUser(nameLower)
}

func (u *userUseCase) GetUserById(id int) (model.User, error) {
	return u.repo.GetUserById(id)
}

func (u *userUseCase) UpdateUser(user model.User, id int) (model.User, error) {
	pastAccount, err := u.GetUserById(id)
	if err != nil {
		return model.User{}, err
	}

	// Cek apakah email sama dengan sebelumnya
	if pastAccount.Email != user.Email {
		exists, err := u.repo.IsUserEmailExists(user.Email)
		if err != nil {
			return model.User{}, err
		}
		if exists {
			return model.User{}, fmt.Errorf("user with Email %s already exists", user.Email)
		}
	}

	if user.Role == "" {
		return model.User{}, fmt.Errorf("role are required")
	}

	if user.Role != "admin" && user.Role != "user" && user.Role != "editor" {
		return model.User{}, fmt.Errorf("choose role between admin, user, or editor")
	}

	if len(user.Password) < 8 {
		return model.User{}, fmt.Errorf("password must be at least 8 characters")
	}

	hashedPassword, err := service.HashPassword(user.Password)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to hash password: %v", err)
	}
	user.Password = hashedPassword

	return u.repo.UpdateUser(user, id)
}

func (u *userUseCase) DeleteUser(id int) error {
	_, err := u.GetUserById(id)
	if err != nil {
		return err
	}

	return u.repo.DeleteUser(id)
}

func (u *userUseCase) IsUserIdExists(id int) (bool, error) {
	return u.repo.IsUserIdExists(id)
}

func (u *userUseCase) IsUserEmailExists(email string) (bool, error) {
	return u.repo.IsUserEmailExists(email)
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
