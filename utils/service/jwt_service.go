package service

import (
	"fmt"
	"time"
	"web-article/config"
	"web-article/model"
	modelutils "web-article/utils/model_utils"

	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	cfg config.TokenConfig
}

type JWTService interface {
	CreateToken(user model.UserLogin) (string, error)
	VerifyToken(tokenString string) (modelutils.JWTPayloadClaim, error)
}

func (j *jwtService) CreateToken(user model.UserLogin) (string, error) {
	tokenKey := j.cfg.JWTSignatureKey

	claims := modelutils.JWTPayloadClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.ApplicationName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.AccessTokenLifeTime)),
		},
		UserId: user.ID,
		Role:   user.Role,
	}

	jwtNewClaim := jwt.NewWithClaims(j.cfg.JWTSigningMethod, claims)

	token, err := jwtNewClaim.SignedString(tokenKey)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (j *jwtService) VerifyToken(tokenString string) (modelutils.JWTPayloadClaim, error) {
	tokenParse, err := jwt.ParseWithClaims(tokenString, &modelutils.JWTPayloadClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return j.cfg.JWTSignatureKey, nil
		},
	)
	if err != nil {
		return modelutils.JWTPayloadClaim{}, err
	}

	claim, ok := tokenParse.Claims.(*modelutils.JWTPayloadClaim)
	if !ok {
		return modelutils.JWTPayloadClaim{}, fmt.Errorf("error claim")
	}

	return *claim, nil
}

func NewJWTService(cfg config.TokenConfig) JWTService {
	return &jwtService{cfg: cfg}
}
