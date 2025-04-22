package router

import (
	"xyz/middlewares"
	"xyz/modules/installment/handler"
	"xyz/modules/installment/repository"
	"xyz/modules/installment/service"

	limitRepository "xyz/modules/limit/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InstallmentRouter(installment *echo.Group, db *gorm.DB) {
	installmentCommandRepository := repository.NewInstallmentCommandRepository(db)
	installmentQueryRepository := repository.NewInstallmentQueryRepository(db)

	limitCommnadRepository := limitRepository.NewLimitCommandRepository(db)

	installmentCommandService := service.NewInstallmentCommandService(installmentCommandRepository, installmentQueryRepository, limitCommnadRepository)
	installmentQueryService := service.NewInstallmentQueryService(installmentQueryRepository)
	installmentHandler := handler.NewInstallmentHandler(installmentQueryService, installmentCommandService)

	// Installment
	installment.PUT("/:installment_id", installmentHandler.UpdateInstallment, middlewares.JWTMiddleware())
	installment.GET("/:installment_id", installmentHandler.GetInstallmentByID, middlewares.JWTMiddleware())
	installment.GET("/transactions/near-due-date/:transaction_id", installmentHandler.GetInstallmentsNearDueDate, middlewares.JWTMiddleware())
	installment.GET("/transactions/:transaction_id", installmentHandler.GetInstallmentsByTransactionID, middlewares.JWTMiddleware())
	installment.GET(("/transactions/near-due-date"), installmentHandler.GetInstallmentsNearDueDateWithoutIds, middlewares.JWTMiddleware())
}