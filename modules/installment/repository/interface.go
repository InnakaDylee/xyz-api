package repository

import "xyz/modules/installment/domain"

type InstallmentCommandRepositoryInterface interface {
	CreateInstallment(intsallment domain.Installment) (domain.Installment, error)
	UpdateInstallment(intsallment domain.Installment) (domain.Installment, error)
}

type InstallmentQueryRepositoryInterface interface {
	GetInstallmentByID(id string) (domain.Installment, error)
	GetInstallmentsByTransactionID(transactionId string) ([]domain.Installment, error)
	GetInstallmentsNearDueDate(transactionId string) ([]domain.Installment, error)
	GetInstallmentsNearDueDateWithoutIds() ([]domain.Installment, error)
}