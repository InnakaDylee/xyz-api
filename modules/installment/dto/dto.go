package dto

import "xyz/modules/installment/domain"

type (
	InstallmentRequest struct {
		Status           string `json:"status" form:"status" validate:"required"`
	}
	
	InstallmentResponse struct {
	ID    string     `json:"installment_id"`
	TransactionID    string     `json:"transaction_id"`
	InstallmentOrder int        `json:"installment_order"`
	PaymentAmount    int        `json:"payment_amount"`
	DueDate          string     `json:"due_date"`
	PaymentDate      *string    `json:"payment_date"`
	Status           string     `json:"status"`
	CreatedAt        string     `json:"created_at"`
	UpdatedAt        string     `json:"updated_at"`
}
)

func InstallmentDomainIntoResponse(installment domain.Installment) InstallmentResponse {
	return InstallmentResponse{
		ID:              installment.ID,
		TransactionID:   installment.TransactionID,
		InstallmentOrder: installment.InstallmentOrder,
		PaymentAmount:   installment.PaymentAmount,
		DueDate:        installment.DueDate.Format("2006-01-02"),
		PaymentDate:     nil,
		Status:         installment.Status,
		CreatedAt:       installment.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:       installment.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func InstallmentDomainIntoResponseList(installments []domain.Installment) []InstallmentResponse {
	var installmentList []InstallmentResponse
	for _, installment := range installments {
		installmentList = append(installmentList, InstallmentDomainIntoResponse(installment))
	}
	return installmentList
}

func InstallmentRequestIntoDomain(request InstallmentRequest, installmentId string) domain.Installment {
	return domain.Installment{
		ID: installmentId,
		Status: request.Status,
	}
}