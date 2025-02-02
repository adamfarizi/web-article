package controller

import (
	"net/http"
	"strconv"
	"web-article/middleware"
	"web-article/model"
	"web-article/usecase"

	"github.com/gin-gonic/gin"
)

type commentController struct {
	useCase        usecase.CommentUsecase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (c *commentController) Route() {
	// Grup route untuk admin
	userAdminRoute := c.rg.Group("/comment", c.authMiddleware.RequireToken("user", "admin"))
	{
		// Example route with middleware
		userAdminRoute.POST("/", c.createComment)
		userAdminRoute.PUT("/:id", c.updateComment)
		userAdminRoute.DELETE("/:id", c.deleteComment)
	}
}

func (cmn *commentController) createComment(c *gin.Context) {
	var payload model.Comment

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment, err := cmn.useCase.CreateComment(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, struct {
		Message string        `json:"message"`
		Data    model.Comment `json:"data"`
	}{
		Message: "Comment created successfully",
		Data:    comment,
	})
}

func (cmn *commentController) updateComment(c *gin.Context) {
	id := c.Param("id")
	commentID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article id"})
		return
	}

	var payload model.Comment

	err = c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment, err := cmn.useCase.UpdateComment(payload, commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, struct {
		Message string        `json:"message"`
		Data    model.Comment `json:"data"`
	}{
		Message: "Comment updated successfully",
		Data:    comment,
	})
}

func (cmn *commentController) deleteComment(c *gin.Context) {
	id := c.Param("id")
	commentID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article id"})
		return
	}

	err = cmn.useCase.DeleteComment(commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

func NewCommentController(useCase usecase.CommentUsecase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *commentController {
	return &commentController{useCase: useCase, rg: rg, authMiddleware: authMiddleware}
}
