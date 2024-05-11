package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// TODO - These functions have not been used in validations
// Figure out how to include them in schema validation
func init() {
	fmt.Println("inside validators init")
	validate := validator.New()
	if err := validate.RegisterValidation("mobile_number", IsValidMobileNumber); err != nil {
		panic(err)
	}
	fmt.Println("success1")
	if err := validate.RegisterValidation("password", IsValidPassword); err != nil {
		panic(err)
	}
	fmt.Println("success2")
	fmt.Println("customer vlaidations - ")
}
