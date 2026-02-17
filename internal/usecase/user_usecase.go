package usecase

import (
	"assignment2/internal/repository"
	"assignment2/pkg/modules"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) GetUsers() ([]modules.User, error) {
	return u.repo.GetUsers()
}

func (u *UserUsecase) GetUserByID(id int) (*modules.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *UserUsecase) CreateUser(name, email string, age int) (int, error) {
	return u.repo.CreateUser(name, email, age)
}

func (u *UserUsecase) UpdateUser(id int, name, email string, age int) (int, error) {
	return u.repo.UpdateUser(id, name, email, age)
}

func (u *UserUsecase) DeleteUser(id int) (int, error) {
	return u.repo.DeleteUser(id)
}