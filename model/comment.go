package model

import "time"

// Comment struct untuk tabel "comments"
type Comment struct {
	ID        int       `json:"id"`                            // Primary key
	ArticleID int       `json:"article_id" binding:"required"` // ID artikel, wajib
	UserID    int       `json:"user_id" binding:"required"`    // ID pengguna, wajib
	Content   string    `json:"content" binding:"required"`    // Isi komentar, wajib
	CreatedAt time.Time `json:"created_at"`                    // Timestamp pembuatan
	UpdatedAt time.Time `json:"updated_at"`                    // Timestamp pembaruan
}

type CommentArticleUser struct {
	ID        int            `json:"id"`
	Article   ArticleComment `json:"article" binding:"required"`
	User      UserComment    `json:"user" binding:"required"`
	Content   string         `json:"content" binding:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
