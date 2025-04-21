package handler

import (
	"github.com/labstack/echo/v4"
)

type ConsumerHandlerInterface interface {
	UpdateConsumer(ctx echo.Context) error
	
	GetConsumerByUserID(ctx echo.Context) error
}