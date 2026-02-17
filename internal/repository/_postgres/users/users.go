package users

import (
	"assignment2/internal/repository/_postgres"
	"assignment2/pkg/modules"
	"errors"
	"time"
)

type Repository struct {
	db *_postgres.Dialect
	executionTimeout time.Duration
}

func NewUserRepository(db *_postgres.Dialect) *Repository {
	return &Repository{
		db:               db,
		executionTimeout: time.Second * 5,
	}
}

func (r *Repository) GetUsers() ([]modules.User, error) {
	var userList []modules.User
	err := r.db.DB.Select(&userList, "SELECT id, name, email, age, created_at FROM users")
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func (r *Repository) GetUserByID(id int) (*modules.User, error) {
	var user modules.User
	err := r.db.DB.Get(&user, "SELECT id, name, email, age, created_at FROM users WHERE id=$1", id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *Repository) CreateUser(name, email string, age int) (int, error) {
	if name == "" {
		return 0, errors.New("name is required")
	}
	if email == "" {
		return 0, errors.New("email is required")
	}
	if age <= 0 {
		return 0, errors.New("age must be positive")
	}

	var id int
	err := r.db.DB.QueryRow(
		"INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id",
		name, email, age,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repository) UpdateUser(id int, name, email string, age int) (int, error) {
	result, err := r.db.DB.Exec(
		"UPDATE users SET name=$1, email=$2, age=$3 WHERE id=$4",
		name, email, age, id,
	)
	if err != nil {
		return 0, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return 0, errors.New("user not found")
	}
	return int(rowsAffected), nil
}

func (r *Repository) DeleteUser(id int) (int, error) {
	result, err := r.db.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return 0, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return 0, errors.New("user not found")
	}
	return int(rowsAffected), nil
}