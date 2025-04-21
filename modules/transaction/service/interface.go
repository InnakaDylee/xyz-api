package service

import "xyz/modules/transaction/domain"

type TransactionCommandServiceInterface interface {
	CreateTransaction(transaction domain.Transaction) (domain.Transaction, error)
}

type TransactionQueryServiceInterface interface {
	GetTransactionByID(transactionID string) (domain.Transaction, error)
	GetTransactionByConsumerID(userID string) ([]domain.Transaction, error)
}