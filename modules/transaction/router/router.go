package router

import (
	"xyz/middlewares"
	consumerRepository "xyz/modules/consumer/repository"
	consumerService "xyz/modules/consumer/service"
	limitRepository "xyz/modules/limit/repository"
	installmentRepository "xyz/modules/installment/repository"
    installmentService "xyz/modules/installment/service"
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
	installmentCRepository := installmentRepository.NewInstallmentCommandRepository(db)
	installmentQRepository := installmentRepository.NewInstallmentQueryRepository(db)

	transactionCommandService := service.NewTransactionCommandService(transactionCRepository, limitQRepository, limitCRepository)
	transactionQueryService := service.NewTransactionQueryService(transactionQRepository)
	consumerQService := consumerService.NewConsumerQueryService(consumerQRepository)
	installmentCService := installmentService.NewInstallmentCommandService(installmentCRepository, installmentQRepository, limitCRepository)

	transactionHandler := handler.NewTransactionHandler(transactionCommandService, transactionQueryService, consumerQService, installmentCService)
	
	transaction.POST("", transactionHandler.CreateTransaction, middlewares.JWTMiddleware())
	transaction.GET("", transactionHandler.GetTransactionByConsumerID, middlewares.JWTMiddleware())
	transaction.GET("/:transaction_id", transactionHandler.GetTransactionByID, middlewares.JWTMiddleware())
}