package handler

import "github.com/labstack/echo/v4"

type InstallmenHendlerInterface interface{
	UpdateInstallment(ctx echo.Context) error
	GetInstallmentByID(ctx echo.Context) error
	GetInstallmentsByTransactionID(ctx echo.Context) error
}