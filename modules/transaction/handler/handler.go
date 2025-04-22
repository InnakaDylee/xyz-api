package handler

import (

	"net/http"
	"xyz/middlewares"
	"xyz/modules/transaction/dto"
	"xyz/modules/transaction/service"
	"xyz/packages/responses"

	consumerService "xyz/modules/consumer/service"
	installmentService "xyz/modules/installment/service"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionCommandService service.TransactionCommandServiceInterface
	transactionQueryService   service.TransactionQueryServiceInterface

	consumerQueryService consumerService.ConsumerQueryServiceInterface

	installmentCommandService installmentService.InstallmentCommandServiceInterface
}

func NewTransactionHandler(
	transactionCommandService service.TransactionCommandServiceInterface,
	transactionQueryService service.TransactionQueryServiceInterface,
	consumerQueryService consumerService.ConsumerQueryServiceInterface,
	installmentCommandService installmentService.InstallmentCommandServiceInterface,
) *TransactionHandler {
	return &TransactionHandler{
		transactionCommandService: transactionCommandService,
		transactionQueryService:   transactionQueryService,
		consumerQueryService:      consumerQueryService,
		installmentCommandService:   installmentCommandService,
	}
}

func (h *TransactionHandler) CreateTransaction(ctx echo.Context) error {
	transaction := dto.TransactionRequest{}
	if err := ctx.Bind(&transaction); err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error()))
	}

	tokenUserId, _, err := middlewares.ExtractToken(ctx)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, responses.ErrorResponse(err.Error()))
	}

	consumer, err := h.consumerQueryService.GetConsumerByUserID(tokenUserId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("Consumer not found"))
	}

	transactionDomain := dto.TransactionRequestIntoDomain(transaction)
	transactionDomain.ConsumerID = consumer.ID


	transactionData, err := h.transactionCommandService.CreateTransaction(transactionDomain)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err.Error()))
	}
	
	transactionDomainData, err := h.transactionQueryService.GetTransactionByID(transactionData.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err.Error()))
	}

	err = h.installmentCommandService.CreateInstallment(transactionDomainData)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err.Error()))
	}

	transactionResponse := dto.TransactionDomainIntoResponse(transactionDomainData)

	return ctx.JSON(http.StatusOK, responses.SuccessResponse("Transaction created successfully", transactionResponse))
}

func (h *TransactionHandler) GetTransactionByID(ctx echo.Context) error {
	transactionID := ctx.Param("transaction_id")
	if transactionID == "" {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("Transaction ID is required"))
	}

	tokenUserId, _, err := middlewares.ExtractToken(ctx)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, responses.ErrorResponse(err.Error()))
	}

	consumer, err := h.consumerQueryService.GetConsumerByUserID(tokenUserId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("Consumer not found"))
	}

	transaction, err := h.transactionQueryService.GetTransactionByID(transactionID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Transaction not found"))
	}

	if transaction.ConsumerID != consumer.ID {
		return ctx.JSON(http.StatusForbidden, responses.ErrorResponse("You are not authorized to access this transaction"))
	}

	transactionResponse := dto.TransactionDomainIntoResponse(transaction)

	return ctx.JSON(http.StatusOK, responses.SuccessResponse("Transaction retrieved successfully", transactionResponse))
}

func (h *TransactionHandler) GetTransactionByConsumerID(ctx echo.Context) error {
	tokenUserId, _, err := middlewares.ExtractToken(ctx)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, responses.ErrorResponse(err.Error()))
	}

	consumer, err := h.consumerQueryService.GetConsumerByUserID(tokenUserId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("Consumer not found"))
	}

	transactions, err := h.transactionQueryService.GetTransactionByConsumerID(consumer.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Transactions not found"))
	}

	transactionResponses := dto.TransactionDomainIntoResponseList(transactions)

	return ctx.JSON(http.StatusOK, responses.SuccessResponse("Transactions retrieved successfully", transactionResponses))
}