package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logger(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, remote_ip=${remote_ip}, host=${host}, protocol=${protocol}, path=${path}, status=${status}, error=${error}\n",
	}))
}