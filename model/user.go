package model

import "time"

// User struct untuk tabel "users"
type User struct {
	ID        int       `json:"id"`                             // Primary key
	Name      string    `json:"name" binding:"required"`        // Nama pengguna, wajib
	Email     string    `json:"email" binding:"required,email"` // Email unik, wajib
	Password  string    `json:"password" binding:"required"`    // Kata sandi, wajib
	Role      string    `json:"role"`                           // Peran, default 'user'
	CreatedAt time.Time `json:"created_at"`                     // Timestamp pembuatan
	UpdatedAt time.Time `json:"updated_at"`                     // Timestamp pembaruan
}

type UserLogin struct {
	ID       int    `json:"id"`                             
	Email    string `json:"email" binding:"required,email"` 
	Password string `json:"password" binding:"required"`    
	Role     string `json:"role"`                           
}

type UserArticle struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserComment struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
