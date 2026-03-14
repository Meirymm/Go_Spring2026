package users

import (
	"assignment4/internal/repository/_postgres"
	"assignment4/pkg/modules"
	"errors"
	"fmt"
	"time"
)

type Repository struct {
	db               *_postgres.Dialect
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
	err := r.db.DB.Select(&userList, "SELECT id, name, email, age, gender, birth_date, created_at FROM users")
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func (r *Repository) GetUserByID(id int) (*modules.User, error) {
	var user modules.User
	err := r.db.DB.Get(&user, "SELECT id, name, email, age, gender, birth_date, created_at FROM users WHERE id=$1", id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *Repository) CreateUser(name, email, gender string, age int, birthDate time.Time) (int, error) {
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
		"INSERT INTO users (name, email, age, gender, birth_date) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		name, email, age, gender, birthDate,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repository) UpdateUser(id int, name, email, gender string, age int, birthDate time.Time) (int, error) {
	result, err := r.db.DB.Exec(
		"UPDATE users SET name=$1, email=$2, age=$3, gender=$4, birth_date=$5 WHERE id=$6",
		name, email, age, gender, birthDate, id,
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

// GetPaginatedUsers - пагинация + фильтрация + сортировка
func (r *Repository) GetPaginatedUsers(page, pageSize int, filters map[string]interface{}, orderBy string) (*modules.PaginatedResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	
	offset := (page - 1) * pageSize
	
	query := "SELECT id, name, email, age, gender, birth_date, created_at FROM users WHERE 1=1"
	countQuery := "SELECT COUNT(*) FROM users WHERE 1=1"
	args := []interface{}{}
	argIndex := 1
	
	// Фильтрация по имени
	if name, ok := filters["name"]; ok && name != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", argIndex)
		countQuery += fmt.Sprintf(" AND name ILIKE $%d", argIndex)
		args = append(args, "%"+name.(string)+"%")
		argIndex++
	}
	
	// Фильтрация по email
	if email, ok := filters["email"]; ok && email != "" {
		query += fmt.Sprintf(" AND email ILIKE $%d", argIndex)
		countQuery += fmt.Sprintf(" AND email ILIKE $%d", argIndex)
		args = append(args, "%"+email.(string)+"%")
		argIndex++
	}
	
	// Фильтрация по gender
	if gender, ok := filters["gender"]; ok && gender != "" {
		query += fmt.Sprintf(" AND gender = $%d", argIndex)
		countQuery += fmt.Sprintf(" AND gender = $%d", argIndex)
		args = append(args, gender)
		argIndex++
	}
	
	// Сортировка
	if orderBy == "" {
		orderBy = "id"
	}
	validColumns := map[string]bool{
		"id": true, "name": true, "email": true, 
		"age": true, "gender": true, "birth_date": true,
	}
	if !validColumns[orderBy] {
		orderBy = "id"
	}
	query += " ORDER BY " + orderBy
	
	// Пагинация
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	paginationArgs := append(args, pageSize, offset)
	
	// Получаем общее количество
	var totalCount int
	err := r.db.DB.Get(&totalCount, countQuery, args...)
	if err != nil {
		return nil, err
	}
	
	// Получаем пользователей
	var users []modules.User
	err = r.db.DB.Select(&users, query, paginationArgs...)
	if err != nil {
		return nil, err
	}
	
	return &modules.PaginatedResponse{
		Data:       users,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
	}, nil
}

// GetCommonFriends - общие друзья двух пользователей
func (r *Repository) GetCommonFriends(userID1, userID2 int) ([]modules.User, error) {
	query := `
		SELECT u.id, u.name, u.email, u.age, u.gender, u.birth_date, u.created_at
		FROM users u
		INNER JOIN user_friends uf1 ON u.id = uf1.friend_id
		INNER JOIN user_friends uf2 ON u.id = uf2.friend_id
		WHERE uf1.user_id = $1 AND uf2.user_id = $2
	`
	
	var users []modules.User
	err := r.db.DB.Select(&users, query, userID1, userID2)
	if err != nil {
		return nil, err
	}
	
	return users, nil
}

// AddFriend - добавить друга
func (r *Repository) AddFriend(userID, friendID int) error {
	if userID == friendID {
		return errors.New("cannot add yourself as friend")
	}
	
	_, err := r.db.DB.Exec(
		"INSERT INTO user_friends (user_id, friend_id) VALUES ($1, $2), ($2, $1) ON CONFLICT DO NOTHING",
		userID, friendID,
	)
	return err
}