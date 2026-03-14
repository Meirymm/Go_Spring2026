package repository

import (
	"assignment4/internal/repository/_postgres"
	"assignment4/internal/repository/_postgres/users"
	"assignment4/pkg/modules"
	"time"
)

type UserRepository interface {
	GetUsers() ([]modules.User, error)
	GetUserByID(id int) (*modules.User, error)
	CreateUser(name, email, gender string, age int, birthDate time.Time) (int, error)
	UpdateUser(id int, name, email, gender string, age int, birthDate time.Time) (int, error)
	DeleteUser(id int) (int, error)
	
	// Новые методы
	GetPaginatedUsers(page, pageSize int, filters map[string]interface{}, orderBy string) (*modules.PaginatedResponse, error)
	GetCommonFriends(userID1, userID2 int) ([]modules.User, error)
	AddFriend(userID, friendID int) error
}

type Repositories struct {
	UserRepository UserRepository
}

func NewRepositories(db *_postgres.Dialect) *Repositories {
	return &Repositories{
		UserRepository: users.NewUserRepository(db),
	}
}