package usecase

import (
	"fmt"
	"strings"
	"web-article/model"
	"web-article/repository"
)

type articleUsecase struct {
	userUseCase UserUseCase
	repo        repository.ArticleRepository
}

type ArticleUsecase interface {
	GetAllArticle(title string) ([]model.ArticleUser, error)
	GetArticleById(id int) (model.ArticleUser, error)
	CreateArticle(article model.Article) (model.Article, error)
	UpdateArticle(article model.Article, id int) (model.Article, error)
	DeleteArticle(articleID int) error
}

func (a *articleUsecase) GetAllArticle(title string) ([]model.ArticleUser, error) {
	titleLower := strings.ToLower(title)
	return a.repo.GetAllArticle(titleLower)
}

func (a *articleUsecase) GetArticleById(id int) (model.ArticleUser, error) {
	return a.repo.GetArticleById(id)
}

func (a *articleUsecase) CreateArticle(article model.Article) (model.Article, error) {
	if article.Title == "" {
		return model.Article{}, fmt.Errorf("tittle are required")
	}

	if article.Content == "" {
		return model.Article{}, fmt.Errorf("content are required")
	}

	_, err := a.userUseCase.GetUserById(article.AuthorID)
	if err != nil {
		return model.Article{}, err
	}

	return a.repo.CreateArticle(article)
}

func (a *articleUsecase) UpdateArticle(article model.Article, id int) (model.Article, error) {
	_, err := a.repo.GetArticleById(id)
	if err != nil {
		return model.Article{}, err
	}

	if article.Title == "" {
		return model.Article{}, fmt.Errorf("tittle are required")
	}

	if article.Content == "" {
		return model.Article{}, fmt.Errorf("content are required")
	}

	_, err = a.userUseCase.GetUserById(article.AuthorID)
	if err != nil {
		return model.Article{}, err
	}

	return a.repo.UpdateArticle(article, id)
}

func (a *articleUsecase) DeleteArticle(articleID int) error {
	_, err := a.repo.GetArticleById(articleID)
	if err != nil {
		return err
	}

	return a.repo.DeleteArticle(articleID)
}

func NewArticleUseCase(uc UserUseCase, repo repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{userUseCase: uc, repo: repo}
}
