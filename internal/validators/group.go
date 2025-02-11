package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// group represents people spending for each other
type Group struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Members     []int     `json:"members" validate:"required,dive,required"`
	Active      bool      `json:"active"`
	CreatedById int       `json:"createdById"`
	CreatedAt   time.Time `json:"createdAt"`
}

func ValidateGroup(group Group) error {
	validate := validator.New()
	return validate.Struct(group)
}
