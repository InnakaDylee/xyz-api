package database

import (
	user "xyz/modules/user/entity"
	consumer "xyz/modules/consumer/entity"
	limit "xyz/modules/limit/entity"
	transaction "xyz/modules/transaction/entity"

	"gorm.io/gorm"
)

func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(
		user.User{},
		consumer.Consumer{},
		limit.Limit{},
		transaction.Transaction{},
	)
}