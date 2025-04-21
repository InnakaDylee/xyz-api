package handler

import "github.com/labstack/echo/v4"

type LimitHandlerInterface interface {
	GetLimitByConsumerId(ctx echo.Context) error
	GetLimitById(ctx echo.Context) error
}