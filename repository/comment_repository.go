package repository

import (
	"database/sql"
	"fmt"
	"web-article/model"
)

type commentRepository struct {
	db *sql.DB
}

type CommentRepository interface {
	GetCommentById(id int) (model.Comment, error)
	CreateComment(comment model.Comment) (model.Comment, error)
	UpdateComment(comment model.Comment, id int) (model.Comment, error)
	DeleteComment(commentID int) error
}

func (c *commentRepository) GetCommentById(id int) (model.Comment, error) {
	var comment model.Comment

	query := "SELECT id, article_id, user_id, content, created_at, updated_at FROM comments WHERE id = $1"
	err := c.db.QueryRow(query, id).Scan(&comment.ID, &comment.ArticleID, &comment.UserID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Comment{}, fmt.Errorf("comment with Id %d not found", id)
		}
		return model.Comment{}, fmt.Errorf("failed to get comment by ID")
	}

	return comment, nil
}

func (c *commentRepository) CreateComment(comment model.Comment) (model.Comment, error) {
	query := "INSERT INTO comments (article_id, user_id, content) VALUES  ($1, $2, $3) RETURNING id, created_at, updated_at;"
	err := c.db.QueryRow(query, comment.ArticleID, comment.UserID, comment.Content).Scan(&comment.ID, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return model.Comment{}, fmt.Errorf("failed to create comment")
	}

	return comment, nil
}

func (c *commentRepository) UpdateComment(comment model.Comment, id int) (model.Comment, error) {
	query := `UPDATE comments SET article_id = $1, user_id = $2, content = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4 RETURNING id, created_at, updated_at`
	err := c.db.QueryRow(query, comment.ArticleID, comment.UserID, comment.Content, id).Scan(&comment.ID, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return model.Comment{}, fmt.Errorf("failed to update comment")
	}

	return comment, nil
}

func (c *commentRepository) DeleteComment(commentID int) error {
	query := `DELETE FROM comments WHERE id = $1`
	_, err := c.db.Exec(query, commentID)
	if err != nil {
		return fmt.Errorf("failed to delete comment")
	}

	return nil
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepository{db: db}
}
