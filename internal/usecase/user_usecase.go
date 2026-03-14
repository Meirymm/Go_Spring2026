package usecase

import (
	"assignment4/internal/repository"
	"assignment4/pkg/modules"
	"time"
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

func (u *UserUsecase) CreateUser(name, email, gender string, age int, birthDate time.Time) (int, error) {
	return u.repo.CreateUser(name, email, gender, age, birthDate)
}

func (u *UserUsecase) UpdateUser(id int, name, email, gender string, age int, birthDate time.Time) (int, error) {
	return u.repo.UpdateUser(id, name, email, gender, age, birthDate)
}

func (u *UserUsecase) DeleteUser(id int) (int, error) {
	return u.repo.DeleteUser(id)
}
func (u *UserUsecase) GetPaginatedUsers(page, pageSize int, filters map[string]interface{}, orderBy string) (*modules.PaginatedResponse, error) {
	return u.repo.GetPaginatedUsers(page, pageSize, filters, orderBy)
}

func (u *UserUsecase) GetCommonFriends(userID1, userID2 int) ([]modules.User, error) {
	return u.repo.GetCommonFriends(userID1, userID2)
}

func (u *UserUsecase) AddFriend(userID, friendID int) error {
	return u.repo.AddFriend(userID, friendID)
}