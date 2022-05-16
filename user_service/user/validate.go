package user

import (
	"github.com/nayefradwi/go_chat_common"
	"net/mail"
)

func checkUserIsValid(user User) *gochatcommon.BaseError {
	validationErrors := make([]gochatcommon.ValidationFieldError, 0)
	// todo: this should be regex?
	if len(user.Username) < 5 || len(user.Username) > 50 {
		validationErrors = append(validationErrors, gochatcommon.NewFieldValidationError("username", "username must be at least 5 charaters"))
	}
	if _, err := mail.ParseAddress(user.Email); err != nil {
		validationErrors = append(validationErrors, gochatcommon.NewFieldValidationError("email", "invalid email format"))
	}
	// todo: this should be regex?
	if len(user.Password) < 8 {
		validationErrors = append(validationErrors, gochatcommon.NewFieldValidationError("password", "password must be at least 8 characters"))
	}
	if len(validationErrors) == 0 {
		return nil
	}
	return gochatcommon.NewValidationError(validationErrors...)
}
