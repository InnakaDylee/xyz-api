package service

import (
	"xyz/modules/installment/domain"
	"xyz/modules/installment/repository"
	"xyz/packages/validator"
	"errors"
)

type InstallmentQueryService struct {
	InstallmentQueryRepository repository.InstallmentQueryRepositoryInterface
}

func NewInstallmentQueryService(installmentQueryRepository repository.InstallmentQueryRepositoryInterface) InstallmentQueryServiceInterface {
	return &InstallmentQueryService{
		InstallmentQueryRepository: installmentQueryRepository,
	}
}

func (s *InstallmentQueryService) GetInstallmentByID(id string) (domain.Installment, error) {
	validateEmpty := validator.CheckEmpty(id)
	if validateEmpty != nil {
		return domain.Installment{}, validateEmpty
	}
	installment, err := s.InstallmentQueryRepository.GetInstallmentByID(id)
	if err != nil {
		return domain.Installment{}, err
	}
	return installment, nil
}

func (s *InstallmentQueryService) GetInstallmentsByTransactionID(transactionId string) ([]domain.Installment, error) {
	validateEmpty := validator.CheckEmpty(transactionId)
	if validateEmpty != nil {
		return nil, validateEmpty
	}
	installments, err := s.InstallmentQueryRepository.GetInstallmentsByTransactionID(transactionId)
	if err != nil {
		return nil, err
	}

	return installments, nil
}

func (s *InstallmentQueryService) GetInstallmentsNearDueDate(transactionId string) ([]domain.Installment, error) {
	validateEmpty := validator.CheckEmpty(transactionId)
	if validateEmpty != nil {
		return nil, validateEmpty
	}
	installments, err := s.InstallmentQueryRepository.GetInstallmentsNearDueDate(transactionId)
	if err != nil {
		return nil, err
	}

	if len(installments) == 0 {
		return nil, errors.New("no installments found")
	}

	return installments, nil
}

func (s *InstallmentQueryService) GetInstallmentsNearDueDateWithoutIds() ([]domain.Installment, error) {
	installments, err := s.InstallmentQueryRepository.GetInstallmentsNearDueDateWithoutIds()
	if err != nil {
		return nil, err
	}

	if len(installments) == 0 {
		return nil, errors.New("no installments found")
	}

	return installments, nil
}