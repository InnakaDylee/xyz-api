package handler

import (
	"net/http"
	"xyz/middlewares"
	"xyz/modules/user/dto"
	"xyz/modules/user/service"
	"xyz/packages/responses"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserCommandService service.UserCommandServiceInterface
	UserQueryService service.UserQueryServiceInterface
}

func NewUserHandler(userCommandService service.UserCommandServiceInterface, userQueryService service.UserQueryServiceInterface) *UserHandler {
	return &UserHandler{
		UserCommandService: userCommandService,
		UserQueryService: userQueryService,
	}
}

func (h *UserHandler) Login(ctx echo.Context) error {
	req := dto.LoginRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error()))
	}
	userDomain := dto.LoginRequestIntoUserDomain(req)
	user, token, err := h.UserCommandService.Login(userDomain.Username, userDomain.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err.Error()))
	}
	response := dto.UserDomainIntoLoginResponse(user, token)

	return ctx.JSON(http.StatusOK, responses.SuccessResponse("Login successful", response))
}
func (h *UserHandler) Register(ctx echo.Context) error {
	req := dto.RegisterRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error()))
	}

	userDomain := dto.RegisterRequestIntoUserDomain(req)

	userRegistered, err := h.UserCommandService.Register(userDomain)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err.Error()))
	}

	response := dto.UserDomainIntoRegisterResponse(userRegistered)

	return ctx.JSON(http.StatusOK, responses.SuccessResponse("User registered successfully", response))
}

func (h *UserHandler) RequestActivation(ctx echo.Context) error {
	req := dto.RequestActivationRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error()))
	}

	userDomain := dto.RequestActivationRequestIntoUserDomain(req)

	err := h.UserCommandService.RequestActivation(userDomain.Email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, responses.SuccessResponse("Activation email sent successfully", nil))
}

func (h *UserHandler) ActivateUser(ctx echo.Context) error {
	reqParam := ctx.QueryParam("token")

	verifyToken, err := middlewares.IsTokenExpired(reqParam)
	if !verifyToken && err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error()))
	}
	email , err := middlewares.ExtractTokenEmail(reqParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error()))
	}

	activateErr := h.UserCommandService.ActivateUser(email)
	if activateErr != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(activateErr.Error()))
	}
	return ctx.JSON(http.StatusOK, responses.SuccessResponse("User activated successfully", nil))
}
func (h *UserHandler) ResetPassword(ctx echo.Context) error {
	return nil

}
func (h *UserHandler) ChangePassword(ctx echo.Context) error {
	return nil

}
func (h *UserHandler) GetUserByID(ctx echo.Context) error {
	return nil

}

