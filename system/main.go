package system

import (
	"baseapp/models"

	"gorm.io/gorm"
)

var (
	SystemAdmin     = models.SystemRole{RoleId: 1, RoleName: "administrator"}
	SystemModerator = models.SystemRole{RoleId: 2, RoleName: "moderator"}
	SystemUser      = models.SystemRole{RoleId: 3, RoleName: "user"}
)

func InitSystemRoles(db *gorm.DB) {
	db.Save(&SystemAdmin)
	db.Save(&SystemModerator)
	db.Save(&SystemUser)
}

func IsAdmin(user *models.User) bool {
	return user.Role.RoleName == SystemAdmin.RoleName
}
