package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"xyz/middlewares"
	"xyz/modules/limit/dto"
	"xyz/modules/limit/service"
	"xyz/packages/responses"

	qService "xyz/modules/consumer/service"

	"github.com/labstack/echo/v4"
)

type LimitHandler struct {
	limitQueryService service.LimitQueryServiceInterface
	
	consumerQueryService qService.ConsumerQueryServiceInterface
}

func NewLimitHandler(limitQueryService service.LimitQueryServiceInterface, consumerQueryService qService.ConsumerQueryServiceInterface) *LimitHandler {
	return &LimitHandler{
		limitQueryService: limitQueryService,
		consumerQueryService: consumerQueryService,
	}
}

func (h *LimitHandler) GetLimitByConsumerId(ctx echo.Context) error {
	userId := ctx.Param("userId")

	tokenUserId, _, err := middlewares.ExtractToken(ctx)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, responses.ErrorResponse(err.Error()))
	}
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println(userId)
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("Invalid userId"))
	}

	if tokenUserId != userIdInt {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("UserId is required"))
	}

	consumerId, err := h.consumerQueryService.GetConsumerByUserID(tokenUserId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err.Error()))
	}

	if consumerId.ID == "" {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ConsumerId is required"))
	}

	limit, err := h.limitQueryService.GetLimitByConsumerId(consumerId.ID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, responses.ErrorResponse(err.Error()))
	}

	limitResponse := dto.LimitDomainIntoListResponse(limit)

	return ctx.JSON(http.StatusOK, responses.SuccessResponse("Limit found", limitResponse))
}

func (h *LimitHandler) GetLimitById(ctx echo.Context) error {
	tokenUserId, _, err := middlewares.ExtractToken(ctx)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, responses.ErrorResponse(err.Error()))
	}

	if tokenUserId == 0 {
		return ctx.JSON(http.StatusUnauthorized, responses.ErrorResponse("Unauthorized"))
	}

	_, err = h.consumerQueryService.GetConsumerByUserID(tokenUserId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err.Error()))
	}

	limitId := ctx.Param("limitId")
	if limitId == "" {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("LimitId is required"))
	}

	limit, err := h.limitQueryService.GetLimitById(limitId)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, responses.ErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, responses.SuccessResponse("Limit found", limit))
}