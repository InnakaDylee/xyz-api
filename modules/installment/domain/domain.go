package domain

import (
	"time"
	"xyz/modules/installment/entity"
	transactionDomain "xyz/modules/transaction/domain"
)

// Installment represents the installment details.
type Installment struct {
	ID              string    `json:"id"`
	TransactionID   string    `json:"transaction_id"`
	Transaction     transactionDomain.Transaction `gorm:"foreignKey:TransactionID;references:ID"`
	InstallmentOrder int      `json:"installment_order"`
	PaymentAmount   int       `json:"payment_amount"`
	DueDate         time.Time `json:"due_date"`
	PaymentDate     *time.Time `json:"payment_date"`
	Status          string    `json:"status"`
	CreatedAt	   time.Time `json:"created_at"`
	UpdatedAt	   time.Time `json:"updated_at"`
}

func InstallmentDomainIntoEntity(installment Installment) entity.Installment {
	return entity.Installment{
		ID:              installment.ID,
		TransactionID:   installment.TransactionID,
		Transaction:     transactionDomain.TransactionDomainIntoEntity(installment.Transaction),
		InstallmentOrder: installment.InstallmentOrder,
		PaymentAmount:   installment.PaymentAmount,
		DueDate:        installment.DueDate,
		PaymentDate:    installment.PaymentDate,
		Status:         installment.Status,
	}
}

func InstallmentEntityIntoDomain(installment entity.Installment) Installment {
	return Installment{
		ID:              installment.ID,
		TransactionID:   installment.TransactionID,
		Transaction:     transactionDomain.TransactionEntityIntoDomain(installment.Transaction),
		InstallmentOrder: installment.InstallmentOrder,
		PaymentAmount:   installment.PaymentAmount,
		DueDate:        installment.DueDate,
		PaymentDate:    installment.PaymentDate,
		Status:         installment.Status,
	}
}

func InstallmentEntityIntoDomainList(installments []entity.Installment) []Installment {
	var installmentList []Installment
	for _, installment := range installments {
		installmentList = append(installmentList, InstallmentEntityIntoDomain(installment))
	}
	return installmentList
}