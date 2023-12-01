package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User holds information relating to users that use the application
type User struct {
	gorm.Model
	Email           string `gorm:"unique"`
	Username        string `gorm:"unique"`
	Password        string
	ActivatedAt     *time.Time
	Tokens          []Token `gorm:"polymorphic:Model;"`
	Sessions        []Session
	UserRoleId      uint
	Role            SystemRole          `gorm:"ForeignKey:UserRoleId"`
	Notifications   []Notification      `gorm:"ForeignKey:UserID"`
	ActiveBans      []Report	        `gorm:"ForeignKey:UserReportedID"`
}

type SystemRole struct {
	RoleId   uint `gorm:"primaryKey"`
	RoleName string
}

// func (user *User) SetPassword(password string) {
// 	user.Password = password
// }

func (user *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

type Notification struct {
    ID                  int         `gorm:"primaryKey"`
    UserID              int
	Type                string
	Content             string
}

type Report struct {
	ID   				uint		`gorm:"primaryKey"`
	PostReportedID		*uint
	UserReportedID		*uint
	UserReportingID		*uint
	AcceptedByID		*uint
	CreatedAt			time.Time	`gorm:"created_at"`
	UpdatedAt			time.Time	`gorm:"updated_at"`
	AcceptedBy		    User	    `gorm:"ForeignKey:AcceptedByID"`
	BanStartsAt			*time.Time	`gorm:"starts_at"`
	BanEndsAt		    *time.Time	`gorm:"ends_at"`
	UserReported		User	    `gorm:"ForeignKey:UserReportedID"`
	UserReporting		User	    `gorm:"ForeignKey:UserReportingID"`
}


