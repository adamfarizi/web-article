package usecase

import (
	"fmt"
	"web-article/model"
	"web-article/utils/service"
)

type authenticationUseCase struct {
	userUseCase UserUseCase
	jwtService  service.JWTService
}

type AuthenticationUseCase interface {
	LoginHandler(email string, password string) (string, error)
	RegisterHandler(user model.User) (model.User, error)
}

func (a *authenticationUseCase) LoginHandler(email string, password string) (string, error) {
	user, err := a.userUseCase.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	// Membandingkan hashed password dari database dengan password yang diterima
	err = service.ComparePassword(user.Password, password)
	if err != nil {
		// Jika password tidak cocok
		return "", fmt.Errorf("password salah")
	}

	token, err := a.jwtService.CreateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *authenticationUseCase) RegisterHandler(user model.User) (model.User, error) {
	return a.userUseCase.CreateUser(user)
}

func NewAuthenticationUsecase(uc UserUseCase, jwtService service.JWTService) AuthenticationUseCase {
	return &authenticationUseCase{userUseCase: uc, jwtService: jwtService}
}
