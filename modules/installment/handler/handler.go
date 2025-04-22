package handler

import (
	"net/http"
	"xyz/modules/installment/dto"
	"xyz/modules/installment/service"
	"xyz/packages/responses"
	"xyz/packages/validator"

	"github.com/labstack/echo/v4"
)

type InstallmentHandler struct {
	installmentQueryService service.InstallmentQueryServiceInterface
	installmentCommandService service.InstallmentCommandServiceInterface
}

func NewInstallmentHandler(
	installmentQueryService service.InstallmentQueryServiceInterface,
	installmentCommandService service.InstallmentCommandServiceInterface,
) InstallmentHandler {
	return InstallmentHandler{
		installmentQueryService:  installmentQueryService,
		installmentCommandService: installmentCommandService,
	}
}

func (h *InstallmentHandler) GetInstallmentByID(ctx echo.Context) (error) {
	id := ctx.Param("installment_id")

	validateEmpty := validator.CheckEmpty(id)
	if validateEmpty != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(validateEmpty.Error()))
	}
	installment, err := h.installmentQueryService.GetInstallmentByID(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error()))
	}

	installmentResponse := dto.InstallmentDomainIntoResponse(installment)

	return ctx.JSON(http.StatusOK, responses.SuccessResponse("Success Get Installment", installmentResponse))
}

func (h *InstallmentHandler) GetInstallmentsByTransactionID(ctx echo.Context) (error) {
	transactionID := ctx.Param("transaction_id")

	validateEmpty := validator.CheckEmpty(transactionID)
	if validateEmpty != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(validateEmpty.Error()))
	}
	installments, err := h.installmentQueryService.GetInstallmentsByTransactionID(transactionID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error()))
	}

	installmentsResponse := dto.InstallmentDomainIntoResponseList(installments)

	return ctx.JSON(http.StatusOK, responses.SuccessResponse("Success Get Installments", installmentsResponse))
}

func (h *InstallmentHandler) GetInstallmentsNearDueDate(ctx echo.Context) (error) {
	transactionID := ctx.Param("transaction_id")

	validateEmpty := validator.CheckEmpty(transactionID)
	if validateEmpty != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(validateEmpty.Error()))
	}
	installments, err := h.installmentQueryService.GetInstallmentsNearDueDate(transactionID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error()))
	}	

	installmentsResponse := dto.InstallmentDomainIntoResponseList(installments)

	return ctx.JSON(http.StatusOK, responses.SuccessResponse("Success Get Installments Near Due Date", installmentsResponse))
}

func (h *InstallmentHandler) GetInstallmentsNearDueDateWithoutIds(ctx echo.Context) (error) {
	installments, err := h.installmentQueryService.GetInstallmentsNearDueDateWithoutIds()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, responses.SuccessResponse("Success Get Installments Near Due Date", installments))
}

func (h *InstallmentHandler) UpdateInstallment(ctx echo.Context) (error) {
	id := ctx.Param("installment_id")

	validateEmpty := validator.CheckEmpty(id)
	if validateEmpty != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(validateEmpty.Error()))
	}

	var installment dto.InstallmentRequest
	if err := ctx.Bind(&installment); err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error()))
	}

	installmentDomain := dto.InstallmentRequestIntoDomain(installment, id)
	updatedInstallment, err := h.installmentCommandService.UpdateInstallment(installmentDomain)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error()))
	}

	updatedInstallmentResponse := dto.InstallmentDomainIntoResponse(updatedInstallment)

	return ctx.JSON(http.StatusOK, responses.SuccessResponse("Success Update Installment", updatedInstallmentResponse))
}