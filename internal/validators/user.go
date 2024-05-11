package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName" validate:"required"`
	LastName  string    `json:"lastName"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Mobile    string    `json:"mobile" validate:"len=10,numeric"`
	Password  string    `json:"password" validate:"required,min=8"`
	CreatedAt time.Time `json:"createdAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

func ValidateUser(user User) error {
	validate := validator.New()
	return validate.Struct(user)
}
