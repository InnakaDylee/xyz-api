package router

import (
	"xyz/middlewares"
	"xyz/modules/limit/handler"
	"xyz/modules/limit/repository"
	"xyz/modules/limit/service"

	consRepository "xyz/modules/consumer/repository"
	consService "xyz/modules/consumer/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func LimitRouter(limit *echo.Group, db *gorm.DB) {
	limitQRepository := repository.NewLimitQueryRepository(db)
	// limitCRepository := repository.NewLimitCommandRepository(db)

	consumerQRepository := consRepository.NewConsumerQueryRepository(db)

	// limitCommandService := service.NewLimitCommandService(limitCRepository, limitQRepository)
	limitQueryService := service.NewLimitQueryService(limitQRepository)

	consumerQueryService := consService.NewConsumerQueryService(consumerQRepository)

	limitHandler := handler.NewLimitHandler(limitQueryService, consumerQueryService)

	limit.GET("/:limitId", limitHandler.GetLimitById, middlewares.JWTMiddleware())
	limit.GET("/consumer/:userId", limitHandler.GetLimitByConsumerId, middlewares.JWTMiddleware())

	limit.GET("/test", func(c echo.Context) error {
		return c.String(200, "Hello, Limit!")
	})
}