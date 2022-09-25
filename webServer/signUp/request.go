package main

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	maxLength       = 256
	namePattern     = `^[a-zA-Z]{2,}`
	emailPattern    = `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`
	passwordPattern = `[ -~]{8,256}$`
)

// global variables to improve performance; local to readable code
// var (
//	nameCheck     = regexp.MustCompile(namePattern)
//	emailCheck    = regexp.MustCompile(emailPattern)
//	passwordCheck = regexp.MustCompile(passwordPattern)
// )

var (
	errInvalidName     = errors.New("nameCheck failed")
	errInvalidEmail    = errors.New("emailCheck failed")
	errInvalidPassword = errors.New("passwordCheck failed")
	errInvalidLength   = errors.New("checkEmail len failed")
)

type InputData struct {
	FullName string `json:"fullName" validate:"min=2"`
	Email    string `json:"email" validate:"max=256,email"`
	Password string `json:"password" validate:"min=8,max=256,ascii"`
}

type UserData struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

func (r *InputData) Validate() error {
	nameCheck := regexp.MustCompile(namePattern)
	emailCheck := regexp.MustCompile(emailPattern)

	if !nameCheck.MatchString(r.FullName) {
		return fmt.Errorf("%w", errInvalidName)
	}

	err := validatePassword(r.Password)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if !emailCheck.MatchString(r.Email) {
		return fmt.Errorf("%w", errInvalidEmail)
	}

	if len(r.Email) > maxLength {
		return fmt.Errorf("%w", errInvalidLength)
	}

	return nil
}

func validatePassword(password string) error {
	passwordCheck := regexp.MustCompile(passwordPattern)

	if !passwordCheck.MatchString(password) {
		return fmt.Errorf("%w", errInvalidPassword)
	}

	return nil
}
