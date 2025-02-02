package controller

import (
	"net/http"
	"web-article/model"
	"web-article/usecase"

	"github.com/gin-gonic/gin"
)

type authController struct {
	authUC usecase.AuthenticationUseCase
	rg     *gin.RouterGroup
}

func (a *authController) Route() {
	a.rg.POST("/auth/login", a.loginHandler)
	a.rg.POST("/auth/register", a.registerHandler)
}

func (a *authController) loginHandler(c *gin.Context) {
	var payload model.UserLogin

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := a.authUC.LoginHandler(payload.Email, payload.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}{
		Message: "Login Success",
		Token:   token,
	})
}

func (a *authController) registerHandler(c *gin.Context) {
	var payload model.User

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := a.authUC.RegisterHandler(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, struct {
		Message string     `json:"message"`
		Data    model.User `json:"data"`
	}{
		Message: "User created successfully",
		Data:    user,
	})
}

func NewAuthController(authUc usecase.AuthenticationUseCase, rg *gin.RouterGroup) *authController {
	return &authController{authUC: authUc, rg: rg}
}
