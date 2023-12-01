package validation

import (
	"baseapp/models"
	"errors"
	"regexp"

	"gorm.io/gorm"
)

type EmailValidation struct {
	email string
	db    *models.Conn
	trans func(s string) string
	Error string
}

func NewEmailValidation(emailParam string, dbParam *models.Conn, transFunc func(s string) string) EmailValidation {
	return EmailValidation{email: emailParam, db: dbParam, trans: transFunc}
}

func (validation *EmailValidation) IsValid() bool {
	if len(validation.email) < 8 {
		validation.Error = validation.trans("Email cannot be shorter than 8 chars")
		return false
	}
	if len(validation.email) > 150 {
		validation.Error = validation.trans("Email cannot be longer than 150 chars")
		return false
	}
	if !validation.isValidEmail() {
		validation.Error = validation.trans("Email not valid")
		return false
	}
	if validation.emailExists() {
		validation.Error = validation.trans("Email is already taken")
		return false
	}
	return true
}

func (validation *EmailValidation) isValidEmail() bool {
	matches, err := regexp.MatchString("^[^@]+@[^@]+\\.[^@]+$", validation.email)
	return err == nil && matches
}

func (validation *EmailValidation) emailExists() bool {
	user := models.User{}
	err := validation.db.Where("email=?", validation.email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
