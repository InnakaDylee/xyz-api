package router

import (
	user "xyz/modules/user/router"
	consumer "xyz/modules/consumer/router"
	limit "xyz/modules/limit/router"
	transaction "xyz/modules/transaction/router"
	installment "xyz/modules/installment/router"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRouter(e *echo.Echo, db *gorm.DB){
	userGroup := e.Group("/api/v1/user")
	user.AuthRouter(userGroup, db)

	consumerGroup := e.Group("/api/v1/consumer")
	consumer.ConsumerRouter(consumerGroup, db)

	limitGroup := e.Group("/api/v1/limit")
	limit.LimitRouter(limitGroup, db)

	transactionGroup := e.Group("/api/v1/transaction")
	transaction.TransactionRouter(transactionGroup, db)

	installmentGroup := e.Group("/api/v1/installment")
	installment.InstallmentRouter(installmentGroup, db)
	
	userGroup.GET("/all", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	},
	)
}