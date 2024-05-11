package validators

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func IsValidPassword(fl validator.FieldLevel) bool {
	fmt.Println("inside isValidPassword")
	password, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	return len(password) >= 8 && regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString(password)
}

func IsValidMobileNumber(fl validator.FieldLevel) bool {
	fmt.Println("inside isValidMobileNumber")
	mobileNumber, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	return regexp.MustCompile(`^\+\d{1,3}\d{6,14}$`).MatchString(mobileNumber)
}
