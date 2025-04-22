package service

import (
	"xyz/modules/installment/domain"
	transactionDomain "xyz/modules/transaction/domain"
)

type InstallmentCommandServiceInterface interface {
	CreateInstallment(transaction transactionDomain.Transaction) error
	UpdateInstallment(installmentInput domain.Installment) (domain.Installment, error)
}

type InstallmentQueryServiceInterface interface {
	GetInstallmentByID(id string) (domain.Installment, error)
	GetInstallmentsByTransactionID(transactionId string) ([]domain.Installment, error)
	GetInstallmentsNearDueDate(transactionId string) ([]domain.Installment, error)
	GetInstallmentsNearDueDateWithoutIds() ([]domain.Installment, error)
}

