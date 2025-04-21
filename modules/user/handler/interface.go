package handler

import "github.com/labstack/echo/v4"

type AuthHandler interface {
	Login(ctx echo.Context) error
	Register(ctx echo.Context) error
	RequestActivation(ctx echo.Context) error
	ActivateUser(ctx echo.Context) error
	ResetPassword(ctx echo.Context) error
	ChangePassword(ctx echo.Context) error
	GetUserByID(ctx echo.Context) error
}