package main

import "errors"

type ValidationError error

var (
	errUsernameExists       = ValidationError(errors.New("username already taken"))
	errEmailExists          = ValidationError(errors.New("an account with that email already exists"))
	errCredentialsIncorrect = ValidationError(errors.New("email, username or password incorrect"))
)

// IsValidationError returns whether the error is of type ValidationError
func IsValidationError(err error) bool {
	_, ok := err.(ValidationError)
	return ok
}
