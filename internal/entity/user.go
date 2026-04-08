package entity
import "github.com/google/uuid"
type User struct {
	ID       uuid.UUID `json:"ID" gorm:"type:uuid;default:uuid_generate_v4()"`
	Username string    `json:"Username"`
	Email    string    `json:"Email"`
	Password string    `json:"Password"`
	Role     string    `json:"Role"`
	Verified bool      `json:"Verified"`}