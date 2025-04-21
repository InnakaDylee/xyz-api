package service

import (
	"xyz/modules/transaction/domain"
	"xyz/modules/transaction/repository"
	"xyz/packages/validator"
)

type TransactionQueryService struct {
	transactionQueryRepostiory repository.TransactionQueryRepositoryInterface
}

func NewTransactionQueryService(transactionQueryRepostiory repository.TransactionQueryRepositoryInterface) TransactionQueryServiceInterface {
	return &TransactionQueryService{
		transactionQueryRepostiory: transactionQueryRepostiory,
	}
}

func (s *TransactionQueryService) GetTransactionByID(transactionID string) (domain.Transaction, error) {
	velidateEmpty := validator.CheckEmpty(transactionID)
	if velidateEmpty != nil {
		return domain.Transaction{}, velidateEmpty
	}

	transaction, err := s.transactionQueryRepostiory.GetTransactionByID(transactionID)
	if err != nil {
		return domain.Transaction{}, err
	}

	return transaction, nil
}

func (s *TransactionQueryService) GetTransactionByConsumerID(consumerID string) ([]domain.Transaction, error) {
	transactions, err := s.transactionQueryRepostiory.GetTransactionByConsumerID(consumerID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}