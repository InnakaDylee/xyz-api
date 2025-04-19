package router

import (
	user "xyz/modules/user/router"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRouter(e *echo.Echo, db *gorm.DB){
	userGroup := e.Group("/api/v1/user")
	user.AuthRouter(userGroup, db)
	
	userGroup.GET("/all", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	},
	)
}