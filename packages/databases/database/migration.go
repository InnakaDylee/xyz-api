package database

import (
	user "xyz/modules/user/entity"

	"gorm.io/gorm"
)

func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(
		user.User{},
	)
}