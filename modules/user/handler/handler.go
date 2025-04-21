package handler

import (
	"net/http"
	"xyz/middlewares"
	consdto "xyz/modules/consumer/dto"
	consService "xyz/modules/consumer/service"
	limit "xyz/modules/limit/domain"
	limitService "xyz/modules/limit/service"
	"xyz/modules/user/dto"
	"xyz/modules/user/service"
	"xyz/packages/responses"
	"xyz/packages/validator"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserCommandService service.UserCommandServiceInterface
	UserQueryService service.UserQueryServiceInterface

	ConsumerCommandService consService.ConsumerCommandServiceInterface
	LimitCommandService limitService.LimitCommandServiceInterface
}

func NewUserHandler(userCommandService service.UserCommandServiceInterface, userQueryService service.UserQueryServiceInterface, consumerCommandService consService.ConsumerCommandServiceInterface, limitCommandService limitService.LimitCommandServiceInterface) *UserHandler {
	return &UserHandler{
		UserCommandService: userCommandService,
		UserQueryService: userQueryService,
		ConsumerCommandService: consumerCommandService,
		LimitCommandService: limitCommandService,
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
	consumerDomain := consdto.RegisterRequestIntoConsumerDomain(req, userDomain.ID)
	
	photoKtp, errKtp := ctx.FormFile("photo_ktp")
	if errKtp != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(errKtp.Error()))
	}
	photoSelfie, errSelfie := ctx.FormFile("photo_selfie")
	if errSelfie != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(errSelfie.Error()))
	}

	isEmpty := validator.CheckEmpty( userDomain.Username, userDomain.Password, userDomain.Email, consumerDomain.Full_Name, consumerDomain.Legal_Name, consumerDomain.Place_Of_Birth, consumerDomain.Date_Of_Birth, consumerDomain.NIK, consumerDomain.Salary, photoKtp, photoSelfie)
	if isEmpty != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(isEmpty.Error()))
	}

	userRegistered, err := h.UserCommandService.Register(userDomain)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("failed to register user"))
	}

	consumerDomain.User_ID = userRegistered.ID
	
	consumerResult, err := h.ConsumerCommandService.CreateConsumer(consumerDomain, photoKtp, photoSelfie)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("failed to create consumer"))
	}

	limit := limit.NewLimit(consumerResult)
	
	_, err = h.LimitCommandService.CreateLimit(limit, consumerResult.Salary)
		if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("failed to create limit"))
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
func (h *UserHandler) ChangePassword(ctx echo.Context) error {
	return nil

}
func (h *UserHandler) GetUserByID(ctx echo.Context) error {
	return nil

}

