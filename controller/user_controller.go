package controller

import (
	"net/http"
	"strconv"
	"web-article/middleware"
	"web-article/model"
	"web-article/usecase"

	"github.com/gin-gonic/gin"
)

type userController struct {
	useCase        usecase.UserUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (u *userController) Route() {
	// Grup route untuk admin
	adminRoutes := u.rg.Group("/users", u.authMiddleware.RequireToken("admin"))
	{
		// Example route with middleware
		adminRoutes.GET("/", u.getAllUser)
		adminRoutes.GET("/:id", u.getUserById)
		adminRoutes.PUT("/:id", u.updateUser)
		adminRoutes.DELETE("/:id", u.deleteUser)
	}
}

func (u *userController) getAllUser(c *gin.Context) {
	name := c.DefaultQuery("name", "")

	users, err := u.useCase.GetAllUser(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found or list is empty"})
		return
	}

	c.JSON(http.StatusOK, struct {
		Message string       `json:"message"`
		Data    []model.User `json:"data"`
	}{
		Message: "User data retrieved successfully",
		Data:    users,
	})

}

func (u *userController) getUserById(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	user, err := u.useCase.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, struct {
		Message string     `json:"message"`
		Data    model.User `json:"data"`
	}{
		Message: "User data retrieved successfully",
		Data:    user,
	})
}

func (u *userController) updateUser(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	var payload model.User

	err = c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.useCase.UpdateUser(payload, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, struct {
		Message string     `json:"message"`
		Data    model.User `json:"data"`
	}{
		Message: "User updated successfully",
		Data:    user,
	})
}

func (u *userController) deleteUser(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	err = u.useCase.DeleteUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func NewUserController(useCase usecase.UserUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *userController {
	return &userController{useCase: useCase, rg: rg, authMiddleware: authMiddleware}
}
