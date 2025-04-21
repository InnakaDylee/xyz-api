package router

import (
	"xyz/middlewares"
	"xyz/modules/consumer/handler"
	"xyz/modules/consumer/repository"
	"xyz/modules/consumer/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ConsumerRouter(consumer *echo.Group, db *gorm.DB) {
	consumerQRepository := repository.NewConsumerQueryRepository(db)
	consumerCRepository := repository.NewConsumerCommandRepository(db)

	consumerCommandService := service.NewConsumerCommandService(consumerCRepository, consumerQRepository)
	consumerQueryService := service.NewConsumerQueryService(consumerQRepository)

	consumerHandler := handler.NewConsumerHandler(consumerCommandService, consumerQueryService)

	consumer.GET("/test", func(c echo.Context) error {
		return c.String(200, "Hello, Consumer!")
	})
	consumer.POST("", consumerHandler.GetConsumerByUserID, middlewares.JWTMiddleware())
}