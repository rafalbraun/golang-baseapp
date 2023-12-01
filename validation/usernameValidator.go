package validation

import (
	"baseapp/models"
	"errors"
	"regexp"

	"gorm.io/gorm"
)

type UsernameValidation struct {
	username string
	db       *models.Conn
	trans    func(s string) string
	Error    string
}

func NewUsernameValidation(usernameParam string, dbParam *models.Conn, transFunc func(s string) string) UsernameValidation {
	return UsernameValidation{username: usernameParam, db: dbParam, trans: transFunc}
}

func (validation *UsernameValidation) IsValid() bool {
	if len(validation.username) < 8 {
		validation.Error = validation.trans("Username cannot be shorter than 8 chars")
		return false
	}
	if len(validation.username) > 30 {
		validation.Error = validation.trans("Username cannot be longer than 30 chars")
		return false
	}
	if !validation.isValidUsername() {
		validation.Error = validation.trans("Username not valid, you can use only letters, digits and underscore")
		return false
	}
	if validation.usernameExists() {
		validation.Error = validation.trans("Username is already taken")
		return false
	}
	return true
}

func (validation *UsernameValidation) isValidUsername() bool {
	matches, err := regexp.MatchString("^[A-Za-z][A-Za-z0-9_]{7,29}$", validation.username)
	return err == nil && matches
}

func (validation *UsernameValidation) usernameExists() bool {
	user := models.User{}
	err := validation.db.Where("username=?", validation.username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
