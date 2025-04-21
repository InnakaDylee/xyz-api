package router

import (
	"xyz/modules/user/handler"
	"xyz/modules/user/repository"
	"xyz/modules/user/service"
	consService "xyz/modules/consumer/service"
	consRepo "xyz/modules/consumer/repository"
	limitService "xyz/modules/limit/service"
	limitRepo "xyz/modules/limit/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AuthRouter(user *echo.Group, db *gorm.DB) {
	// Initialize repositories
	userQRepository := repository.NewUserQueryRepository(db)
	userCRepository := repository.NewUserCommandRepository(db)
	consumerQueryRepository := consRepo.NewConsumerQueryRepository(db)
	consumerCommandRepository := consRepo.NewConsumerCommandRepository(db)
	limitCommandRepository := limitRepo.NewLimitCommandRepository(db)
	limitQueryRepository := limitRepo.NewLimitQueryRepository(db)


	// Initialize services
	userCommandService := service.NewUserCommandService(userCRepository, userQRepository)
	userQueryService := service.NewUserQueryService(userQRepository)
	consumerCommandService := consService.NewConsumerCommandService(consumerCommandRepository, consumerQueryRepository)
	limitCommandService := limitService.NewLimitCommandService(limitCommandRepository, limitQueryRepository)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userCommandService, userQueryService, consumerCommandService, limitCommandService)
	// userHandler := handler.NewUserHandler(userCommandService, userQueryService)


	user.POST("/register", userHandler.Register)
	user.POST("/login", userHandler.Login)
	user.POST("/request-activation", userHandler.RequestActivation)
	user.GET("/activate", userHandler.ActivateUser)
	// user.GET("/:id", )

}