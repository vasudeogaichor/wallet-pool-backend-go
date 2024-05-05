package models

import (
	"time"

	"github.com/go-playground/validator"
)

// Question refers to a discussion for a spending when someone disagrees with it
// NOTE - In future might change to discussions/threads, as it is more appropriate because
// people can have more conversation
type Question struct {
	ID           int       `json:"id"`
	SpendingID   int       `json:"spendingId" validate:"required"`
	QuestionText string    `json:"questionText" validate:"required"`
	Status       string    `json:"status" validate:"required"`
	AskedBy      int       `json:"askedBy" validate:"required"`
	AnsweredBy   int       `json:"answeredBy"`
	Answer       string    `json:"answer"`
	CreatedAt    time.Time `json:"createdAt" validate:"required"`
}

func ValidateQuestion(question Question) error {
	validate := validator.New()
	return validate.Struct(question)
}
