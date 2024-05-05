package models

import (
	"time"

	"github.com/go-playground/validator"
)

// Event represents an event such as a movie outing
type Event struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	GroupID   int       `json:"groupId" validate:"required"`
	Date      time.Time `json:"date" validate:"required"`
	Members   []int     `json:"members"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
}

// ValidateEvent validates the event struct
func ValidateEvent(event Event) error {
	validate := validator.New()
	return validate.Struct(event)
}
