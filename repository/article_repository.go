package repository

import (
	"database/sql"
	"fmt"
	"web-article/model"
)

type articleRepository struct {
	db *sql.DB
}

type ArticleRepository interface {
	GetAllArticle(title string) ([]model.ArticleUser, error)
	GetArticleById(id int) (model.ArticleUser, error)
	CreateArticle(article model.Article) (model.Article, error)
	UpdateArticle(article model.Article, id int) (model.Article, error)
	DeleteArticle(articleID int) error
}

func (a *articleRepository) GetAllArticle(title string) ([]model.ArticleUser, error) {
	var articles []model.ArticleUser

	query := `
		SELECT 
			articles.id, articles.title, articles.content, 
			users.id AS author_id, users.name AS author_name, 
			users.email AS author_email, users.role AS author_role, 
			articles.created_at, articles.updated_at
		FROM articles 
		JOIN users ON articles.author_id = users.id
	`
	var rows *sql.Rows
	var err error

	if title != "" {
		query += " WHERE LOWER(articles.title) LIKE '%' || $1 || '%' ORDER BY articles.id;"
		rows, err = a.db.Query(query, title)
	} else {
		query += " ORDER BY articles.id;"
		rows, err = a.db.Query(query)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve list article")
	}

	for rows.Next() {
		var article model.ArticleUser

		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.Author.ID, &article.Author.Name, &article.Author.Email, &article.Author.Role, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scaning data")
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func (a *articleRepository) GetArticleById(id int) (model.ArticleUser, error) {
	var article model.ArticleUser

	query := `
		SELECT 
			articles.id, articles.title, articles.content, 
			users.id AS author_id, users.name AS author_name, 
			users.email AS author_email, users.role AS author_role, 
			articles.created_at, articles.updated_at
		FROM articles 
		JOIN users ON articles.author_id = users.id
		WHERE articles.id = $1;
	`
	err := a.db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Content, &article.Author.ID, &article.Author.Name, &article.Author.Email, &article.Author.Role, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.ArticleUser{}, fmt.Errorf("article with Id %d not found", id)
		}
		return model.ArticleUser{}, fmt.Errorf("failed to get article by ID")
	}

	return article, nil
}

func (a *articleRepository) CreateArticle(article model.Article) (model.Article, error) {
	query := "INSERT INTO articles (title, content, author_id) VALUES  ($1, $2, $3) RETURNING id, created_at, updated_at"
	err := a.db.QueryRow(query, article.Title, article.Content, article.AuthorID).Scan(&article.ID, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		return model.Article{}, fmt.Errorf("failed to create article")
	}

	return article, nil
}

func (a *articleRepository) UpdateArticle(article model.Article, id int) (model.Article, error) {
	query := `UPDATE articles SET title = $1, content = $2, author_id = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4 RETURNING id, created_at, updated_at`
	err := a.db.QueryRow(query, article.Title, article.Content, article.AuthorID, id).Scan(&article.ID, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		return model.Article{}, fmt.Errorf("failed to update article")
	}

	return article, nil
}

func (a *articleRepository) DeleteArticle(articleID int) error {
	query := `DELETE FROM articles WHERE id = $1`
	_, err := a.db.Exec(query, articleID)
	if err != nil {
		return fmt.Errorf("failed to delete article")
	}

	return nil
}

func NewArticleRepository(db *sql.DB) ArticleRepository {
	return &articleRepository{db: db}
}
