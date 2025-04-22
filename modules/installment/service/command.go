package service

import (
	"xyz/modules/installment/domain"
	"xyz/modules/installment/repository"
	transactionDomain "xyz/modules/transaction/domain"

	limitRepository "xyz/modules/limit/repository"
)

type InstallmentCommandService struct {
	installmentCommandRepository repository.InstallmentCommandRepositoryInterface
	installmentQueryRepository repository.InstallmentQueryRepositoryInterface

	limitCommandRepository limitRepository.LimitCommandRepositoryInterface
}

func NewInstallmentCommandService(
	installmentCommandRepository repository.InstallmentCommandRepositoryInterface,
	installmentQueryRepository repository.InstallmentQueryRepositoryInterface,
	limitCommandRepository limitRepository.LimitCommandRepositoryInterface,
) InstallmentCommandServiceInterface {
	return &InstallmentCommandService{
		installmentCommandRepository: installmentCommandRepository,
		installmentQueryRepository: installmentQueryRepository,
		limitCommandRepository: limitCommandRepository,
	}
}

func (s *InstallmentCommandService) CreateInstallment(transaction transactionDomain.Transaction) (error) {
	installmentInput := domain.Installment{
		TransactionID:   transaction.ID,
		InstallmentOrder: 0,
		PaymentAmount:    int(transaction.InstallmentAmount),
		DueDate:          transaction.TransactionDate,
		PaymentDate:      nil,
		Status:           transaction.Status,
	}


	for i := 1; i <= transaction.Limit.Tenor; i++ {
		installmentInput.InstallmentOrder = i
		installmentInput.DueDate = transaction.TransactionDate.AddDate(0, i, 0)
		_, err := s.installmentCommandRepository.CreateInstallment(installmentInput)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *InstallmentCommandService) UpdateInstallment(installmentInput domain.Installment) (domain.Installment, error) {
	if installmentInput.Status != "paid"{
		return domain.Installment{}, nil
	}

	checkInstallment, err := s.installmentQueryRepository.GetInstallmentByID(installmentInput.ID)
	if err != nil {
		return domain.Installment{}, err
	}
	checkInstallment.Status = installmentInput.Status

	if checkInstallment.Transaction.Limit.Tenor == checkInstallment.InstallmentOrder {
		checkInstallment.Transaction.Limit.RemainingAmount += int(checkInstallment.Transaction.OTR)  
		limit, err := s.limitCommandRepository.UpdateLimit(checkInstallment.Transaction.Limit)
		if err != nil {
			return domain.Installment{}, err
		}
		checkInstallment.Transaction.Limit = limit
	}

	installment, err := s.installmentCommandRepository.UpdateInstallment(checkInstallment)
	if err != nil {
		return domain.Installment{}, err
	}

	return installment, nil
}