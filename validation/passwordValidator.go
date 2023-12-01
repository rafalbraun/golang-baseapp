package validation

import (
	"baseapp/models"
)

type PasswordValidation struct {
	password string
	db       *models.Conn
	trans    func(s string) string
	Error    string
}

func NewPasswordValidation(passwordParam string, dbParam *models.Conn, transFunc func(s string) string) PasswordValidation {
	return PasswordValidation{password: passwordParam, db: dbParam, trans: transFunc}
}

func (validation *PasswordValidation) IsValid() bool {
	if len(validation.password) < 8 {
		validation.Error = validation.trans("Your password must be 8 characters in length or longer")
		return false
	}
	if len(validation.password) > 50 {
		validation.Error = validation.trans("Your password cannot be longer than 50 characters")
		return false
	}
	return true
}
