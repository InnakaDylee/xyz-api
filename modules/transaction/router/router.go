package router

import (
	"xyz/middlewares"
	consumerRepository "xyz/modules/consumer/repository"
	consumerService "xyz/modules/consumer/service"
	limitRepository "xyz/modules/limit/repository"
	"xyz/modules/transaction/handler"
	"xyz/modules/transaction/repository"
	"xyz/modules/transaction/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TransactionRouter(transaction *echo.Group, db *gorm.DB) {
	transactionQRepository := repository.NewTransactionQueryRepository(db)
	transactionCRepository := repository.NewTransactionCommandRepository(db)
	limitQRepository := limitRepository.NewLimitQueryRepository(db)
	limitCRepository := limitRepository.NewLimitCommandRepository(db)
	consumerQRepository := consumerRepository.NewConsumerQueryRepository(db)

	consumerQService := consumerService.NewConsumerQueryService(consumerQRepository)

	transactionCommandService := service.NewTransactionCommandService(transactionCRepository, limitQRepository, limitCRepository)
	transactionQueryService := service.NewTransactionQueryService(transactionQRepository)

	transactionHandler := handler.NewTransactionHandler(transactionCommandService, transactionQueryService, consumerQService)
	
	transaction.POST("", transactionHandler.CreateTransaction, middlewares.JWTMiddleware())
	transaction.GET("", transactionHandler.GetTransactionByConsumerID, middlewares.JWTMiddleware())
	transaction.GET("/:transaction_id", transactionHandler.GetTransactionByID, middlewares.JWTMiddleware())
}