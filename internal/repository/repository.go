package repository

import (
	"assignment2/internal/repository/_postgres"
	"assignment2/internal/repository/_postgres/users"
	"assignment2/pkg/modules"
)

type UserRepository interface {
	GetUsers() ([]modules.User, error)
	GetUserByID(id int) (*modules.User, error)
	CreateUser(name, email string, age int) (int, error)
	UpdateUser(id int, name, email string, age int) (int, error)
	DeleteUser(id int) (int, error)
}

type Repositories struct {
	UserRepository UserRepository
}

func NewRepositories(db *_postgres.Dialect) *Repositories {
	return &Repositories{
		UserRepository: users.NewUserRepository(db),
	}
}