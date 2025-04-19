package router

import (
	"xyz/modules/user/handler"
	"xyz/modules/user/repository"
	"xyz/modules/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AuthRouter(user *echo.Group, db *gorm.DB) {
	// Initialize repositories
	userQRepository := repository.NewUserQueryRepository(db)
	userCRepository := repository.NewUserCommandRepository(db)

	// Initialize services
	userCommandService := service.NewUserCommandService(userCRepository, userQRepository)
	// userQueryService := service.NewUserQueryService(userQRepository)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userCommandService, userQRepository)
	// userHandler := handler.NewUserHandler(userCommandService, userQueryService)


	user.POST("/register", userHandler.Register)
	user.POST("/login", userHandler.Login)
	user.POST("/request-activation", userHandler.RequestActivation)
	user.GET("/activate", userHandler.ActivateUser)
	// user.GET("/:id", )

}