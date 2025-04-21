package repository

import "xyz/modules/transaction/domain"

type TransactionCommandRepositoryInterface interface {
	CreateTransaction(transaction domain.Transaction) (domain.Transaction, error)
}

type TransactionQueryRepositoryInterface interface {
	GetTransactionByID(transactionID string) (domain.Transaction, error)
	GetTransactionByConsumerID(userID string) ([]domain.Transaction, error)
}