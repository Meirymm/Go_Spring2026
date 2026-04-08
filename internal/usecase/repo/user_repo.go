package repo
import (
	"fmt"
	"practice7/internal/entity"
	"practice7/pkg/postgres"
	"github.com/google/uuid")
type UserRepo struct {
	PG *postgres.Postgres}
func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{PG: pg}}
func (u *UserRepo) RegisterUser(user *entity.User) (*entity.User, error) {
	if err := u.PG.Conn.Create(user).Error; err != nil {
		return nil, err}
	return user, nil}

func (u *UserRepo) LoginUser(user *entity.LoginUserDTO) (*entity.User, error) {
	var userFromDB entity.User
	if err := u.PG.Conn.First(&userFromDB, "username = ?", user.Username).Error; err != nil {
		return nil, fmt.Errorf("username not found: %w", err)
	}
	return &userFromDB, nil}

func (u *UserRepo) GetMe(userID uuid.UUID) (*entity.User, error) {
	var user entity.User
	if err := u.PG.Conn.First(&user, "id = ?", userID).Error; err != nil {
		return nil, fmt.Errorf("user not found: %w", err)	}
	return &user, nil}
func (u *UserRepo) PromoteUser(userID uuid.UUID) (*entity.User, error) {
	var user entity.User
	if err := u.PG.Conn.First(&user, "id = ?", userID).Error; err != nil {
		return nil, fmt.Errorf("user not found: %w", err)	}
	user.Role = "admin"
	if err := u.PG.Conn.Save(&user).Error; err != nil {
		return nil, fmt.Errorf("promote failed: %w", err)}
	return &user, nil
}