package handler

import (
	"github.com/labstack/echo/v4"
)

type TransactionHandlerInterface interface {
	CreateTransaction(ctx echo.Context)  error
	GetTransactionByID(ctx echo.Context) error
	GetTransactionByConsumerID(ctx echo.Context) error
}