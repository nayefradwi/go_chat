package user

import (
	"github.com/nayefradwi/go_chat/common/errorHandling"
	"net/mail"
)

func checkUserIsValid(user User) *errorHandling.BaseError {
	validationErrors := make([]errorHandling.ValidationFieldError, 0)
	// todo: this should be regex?
	if len(user.Username) < 5 || len(user.Username) > 50 {
		validationErrors = append(validationErrors, errorHandling.NewFieldValidationError("username", "username must be at least 5 charaters"))
	}
	if _, err := mail.ParseAddress(user.Email); err != nil {
		validationErrors = append(validationErrors, errorHandling.NewFieldValidationError("email", "invalid email format"))
	}
	// todo: this should be regex?
	if len(user.Password) < 8 {
		validationErrors = append(validationErrors, errorHandling.NewFieldValidationError("password", "password must be at least 8 characters"))
	}
	if len(validationErrors) == 0 {
		return nil
	}
	return errorHandling.NewValidationError(validationErrors...)
}
