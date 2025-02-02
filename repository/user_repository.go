package repository

import (
	"database/sql"
	"fmt"
	"web-article/model"
)

type userRepository struct {
	db *sql.DB
}

type UserRepository interface {
	GetUserByEmail(email string) (model.UserLogin, error)
	CreateUser(user model.User) (model.User, error)
	GetAllUser(name string) ([]model.User, error)
	GetUserById(id int) (model.User, error)
	UpdateUser(user model.User, id int) (model.User, error)
	DeleteUser(userID int) error
	IsUserIdExists(id int) (bool, error)
	IsUserEmailExists(email string) (bool, error)
}

func (u *userRepository) GetUserByEmail(email string) (model.UserLogin, error) {
	var user model.UserLogin

	query := "SELECT id, email, password, role FROM users WHERE email = $1 ORDER BY id"
	err := u.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.UserLogin{}, fmt.Errorf("user with Email %s not found", email)
		}
		return model.UserLogin{}, fmt.Errorf("failed to retrieve user by Email")
	}

	return user, nil
}

func (u *userRepository) CreateUser(user model.User) (model.User, error) {
	query := "INSERT INTO users (name, email, password, role) VALUES  ($1, $2, $3, $4) RETURNING id, created_at, updated_at"
	err := u.db.QueryRow(query, user.Name, user.Email, user.Password, user.Role).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to create user")
	}

	return user, nil
}

func (u *userRepository) GetAllUser(name string) ([]model.User, error) {
	var users []model.User

	query := "SELECT id, name, email, role, created_at, updated_at FROM users"

	var rows *sql.Rows
	var err error

	if name != "" {
		query += " WHERE LOWER(name) LIKE '%' || $1 || '%' ORDER BY id;"
		rows, err = u.db.Query(query, name)
	} else {
		query += " ORDER BY id;"
		rows, err = u.db.Query(query)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve list user")
	}

	for rows.Next() {
		var user model.User

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scaning data")
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *userRepository) GetUserById(id int) (model.User, error) {
	var user model.User

	query := "SELECT id, name, email, role, created_at, updated_at FROM users WHERE id = $1"
	err := u.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, fmt.Errorf("user with Id %d not found", id)
		}
		return model.User{}, fmt.Errorf("failed to get user by ID")
	}

	return user, nil
}

func (u *userRepository) UpdateUser(user model.User, id int) (model.User, error) {
	query := `UPDATE users SET name = $1, email = $2, password = $3, role = $4, updated_at = CURRENT_TIMESTAMP WHERE id = $5 RETURNING id, created_at, updated_at`
	err := u.db.QueryRow(query, user.Name, user.Email, user.Password, user.Role, id).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to update user")
	}

	return user, nil
}

func (u *userRepository) DeleteUser(userID int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := u.db.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user")
	}

	return nil
}

func (u *userRepository) IsUserIdExists(id int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)"
	err := u.db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u *userRepository) IsUserEmailExists(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)"
	err := u.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
