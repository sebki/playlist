package errors

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

// ValidationError holds all errors concerning user validation
type ValidationError error

var (
	ErrUsernameExists       = ValidationError(errors.New("username already taken"))
	ErrEmailExists          = ValidationError(errors.New("an account with that email already exists"))
	ErrCredentialsIncorrect = ValidationError(errors.New("email, username or password incorrect"))
)

// IsValidationError returns whether the error is of type ValidationError
func IsValidationError(err error) bool {
	_, ok := err.(ValidationError)
	return ok
}

type ErrorData struct {
	ErrCode int
	ErrMsg  error
}

func InternalServerError(c *gin.Context, err error) {
	errorData := ErrorData{
		ErrCode: 500,
		ErrMsg:  err,
	}

	log.Println(err)
	errorData.ErrMsg = err

	c.JSON(500, errorData)
}
