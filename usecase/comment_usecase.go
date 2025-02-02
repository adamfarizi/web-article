package usecase

import (
	"fmt"
	"web-article/model"
	"web-article/repository"
)

type commentUsecase struct {
	articleUsecase ArticleUsecase
	userUseCase UserUseCase
	repo repository.CommentRepository
}

type CommentUsecase interface {
	CreateComment(comment model.Comment) (model.Comment, error)
	UpdateComment(comment model.Comment, id int) (model.Comment, error)
	DeleteComment(commentID int) error
}

func (c *commentUsecase) CreateComment(comment model.Comment) (model.Comment, error) {
	_, err := c.articleUsecase.GetArticleById(comment.ArticleID)
	if err != nil {
		return model.Comment{}, err
	}

	_, err = c.userUseCase.GetUserById(comment.UserID)
	if err != nil {
		return model.Comment{}, err
	}
	
	if comment.Content == "" {
		return model.Comment{}, fmt.Errorf("content are required")
	}

	return c.repo.CreateComment(comment)
}

func (c *commentUsecase) UpdateComment(comment model.Comment, id int) (model.Comment, error) {
	_, err := c.repo.GetCommentById(id)
	if err != nil {
		return model.Comment{}, err
	}

	_, err = c.articleUsecase.GetArticleById(comment.ArticleID)
	if err != nil {
		return model.Comment{}, err
	}

	_, err = c.userUseCase.GetUserById(comment.UserID)
	if err != nil {
		return model.Comment{}, err
	}
	
	if comment.Content == "" {
		return model.Comment{}, fmt.Errorf("content are required")
	}

	return c.repo.UpdateComment(comment, id)
}

func (c *commentUsecase) DeleteComment(commentID int) error {
	_, err := c.repo.GetCommentById(commentID)
	if err != nil {
		return err
	}

	return c.repo.DeleteComment(commentID)
}

func NewCommentUseCase(repo repository.CommentRepository, articleUsecase ArticleUsecase, userUseCase UserUseCase) CommentUsecase {
	return &commentUsecase{repo: repo, articleUsecase: articleUsecase, userUseCase: userUseCase}
}
