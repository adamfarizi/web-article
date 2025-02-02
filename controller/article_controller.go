package controller

import (
	"net/http"
	"strconv"
	"web-article/middleware"
	"web-article/model"
	"web-article/usecase"

	"github.com/gin-gonic/gin"
)

type articleController struct {
	useCase        usecase.ArticleUsecase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (a *articleController) Route() {

	a.rg.GET("/article", a.getAllUser)

	// Grup route untuk admin
	editorAdminRoute := a.rg.Group("/article", a.authMiddleware.RequireToken("editor", "admin"))
	{
		// Example route with middleware
		editorAdminRoute.POST("/", a.createArticle)
		editorAdminRoute.PUT("/:id", a.updateArticle)
		editorAdminRoute.DELETE("/:id", a.deleteArticle)
	}
}

func (a *articleController) getAllUser(c *gin.Context) {
	title := c.DefaultQuery("title", "")

	articles, err := a.useCase.GetAllArticle(title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(articles) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Article not found or list is empty"})
		return
	}

	c.JSON(http.StatusOK, struct {
		Message string              `json:"message"`
		Data    []model.ArticleUser `json:"data"`
	}{
		Message: "Article data retrieved successfully",
		Data:    articles,
	})

}

func (a *articleController) createArticle(c *gin.Context) {
	var payload model.Article

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := a.useCase.CreateArticle(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, struct {
		Message string        `json:"message"`
		Data    model.Article `json:"data"`
	}{
		Message: "Article created successfully",
		Data:    article,
	})
}

func (a *articleController) updateArticle(c *gin.Context) {
	id := c.Param("id")
	articleID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article id"})
		return
	}

	var payload model.Article

	err = c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := a.useCase.UpdateArticle(payload, articleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, struct {
		Message string        `json:"message"`
		Data    model.Article `json:"data"`
	}{
		Message: "Article updated successfully",
		Data:    article,
	})
}

func (a *articleController) deleteArticle(c *gin.Context) {
	id := c.Param("id")
	articleID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article id"})
		return
	}

	err = a.useCase.DeleteArticle(articleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}

func NewArticleController(useCase usecase.ArticleUsecase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *articleController {
	return &articleController{useCase: useCase, rg: rg, authMiddleware: authMiddleware}
}
