package model

import (
	"time"
)

// Article struct untuk tabel "articles"
type Article struct {
	ID        int       `json:"id"`                           // Primary key
	Title     string    `json:"title" binding:"required"`     // Judul artikel, wajib
	Content   string    `json:"content" binding:"required"`   // Isi artikel, wajib
	AuthorID  int       `json:"author_id" binding:"required"` // ID penulis, wajib
	CreatedAt time.Time `json:"created_at"`                   // Timestamp pembuatan
	UpdatedAt time.Time `json:"updated_at"`                   // Timestamp pembaruan
}

type ArticleUser struct {
	ID        int         `json:"id"`
	Title     string      `json:"title" binding:"required"`
	Content   string      `json:"content" binding:"required"`
	Author    UserArticle `json:"author" binding:"required"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type ArticleComment struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	AuthorID  int       `json:"author_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
