package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Spending represents the instance when someone in the group spent any amount for the group
type Spending struct {
	ID          int       `json:"id"`
	Description string    `json:"description" validate:"required"`
	Amount      float64   `json:"amount" validate:"required,numeric,gt=0"`
	GroupID     int       `json:"groupId"`
	EventID     int       `json:"eventId" validate:"required"`
	Status      string    `json:"status" validate:"required"`
	CreatedBy   int       `json:"createdBy" validate:"required"`
	CreatedAt   time.Time `json:"createdAt"`
}

func ValidateSpending(spending Spending) error {
	validate := validator.New()
	return validate.Struct(spending)
}
