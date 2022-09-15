package main

import (
	"fmt"
	"regexp"
	//_ "github.com/go-playground/validator"
)

const (
	maxLength       = 256
	namePattern     = `^[a-zA-Z]{2,}`
	emailPattern    = `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`
	passwordPattern = `[ -~]{8,256}$`
)

var (
	checkName     = regexp.MustCompile(namePattern)
	checkEmail    = regexp.MustCompile(emailPattern)
	checkPassword = regexp.MustCompile(passwordPattern)
)

type Data struct {
	FullName string `json:"fullName" validate:"min=2"`
	Email    string `json:"email" validate:"max=256,email"` // regular
	Password string `json:"password" validate:"min=8,max=256,ascii"`
}

type userID struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

func (request *Data) Validate() error {
	if !checkName.MatchString(request.FullName) {
		return fmt.Errorf("checkName failed")
	}

	err := validatePassword(request.Password)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if !checkEmail.MatchString(request.Email) {
		return fmt.Errorf("checkEmail failed")
	}

	if len(request.Email) > maxLength {
		return fmt.Errorf("checkEmail len failed")
	}

	return nil
}

func validatePassword(password string) error {
	if !checkPassword.MatchString(password) {
		return fmt.Errorf("checkPassword failed")
	}

	return nil
}

//type Req struct {
//	Data Data `json:"data" validate:"required,dive,required"`
//}
//
//func (r *Req) Validate() error {
//	validate := validator.New()
//
//	if err := validate.Struct(r); err != nil {
//		return fmt.Errorf("%w: Validate failed", err)
//	}
//
//	return nil
//}
