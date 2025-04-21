package handler

import (
	"xyz/middlewares"
	"xyz/modules/consumer/dto"
	"xyz/modules/consumer/service"
	"xyz/packages/responses"

	"github.com/labstack/echo/v4"
)

type ConsumerHandler struct {
	ConsumerCommandService service.ConsumerCommandServiceInterface
	ConsumerQueryService   service.ConsumerQueryServiceInterface
}

func NewConsumerHandler(consumerCommandService service.ConsumerCommandServiceInterface, consumerQueryService service.ConsumerQueryServiceInterface) *ConsumerHandler {
	return &ConsumerHandler{
		ConsumerCommandService: consumerCommandService,
		ConsumerQueryService:   consumerQueryService,
	}
}

func (h *ConsumerHandler) UpdateConsumer(ctx echo.Context) error {
	return nil
}

func (h *ConsumerHandler) GetConsumerByUserID(ctx echo.Context) error {
	UserId, _, err := middlewares.ExtractToken(ctx)
	if err != nil {
		return ctx.JSON(401, responses.ErrorResponse(err.Error()))
	}

	consumer, err := h.ConsumerQueryService.GetConsumerByUserID(UserId)
	if err != nil {
		return ctx.JSON(500, responses.ErrorResponse(err.Error()))
	}

	if consumer.ID == "" {
		return ctx.JSON(404, responses.ErrorResponse("Consumer not found"))
	}

	consumerResponse := dto.ConsumerDomainIntoGetResponse(consumer)

	return ctx.JSON(200, responses.SuccessResponse("Consumer found", consumerResponse))
}