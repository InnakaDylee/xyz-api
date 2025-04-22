package repository

import (
	"xyz/modules/transaction/domain"

	"gorm.io/gorm"
)

type TransactionCommandReposiorty struct {
	db *gorm.DB
}

func NewTransactionCommandRepository(db *gorm.DB) TransactionCommandRepositoryInterface {
	return &TransactionCommandReposiorty{
		db: db,
	}
}

func (r *TransactionCommandReposiorty) CreateTransaction(transactionInput domain.Transaction) (domain.Transaction, error) {
	transaction := r.db.Begin()
	if transaction.Error != nil {
		return domain.Transaction{}, transaction.Error
	}

	transactionRaw := domain.TransactionDomainIntoEntity(transactionInput)
	if err := transaction.Debug().Create(&transactionRaw).Error; err != nil {
		transaction.Rollback()
		return domain.Transaction{}, err
	}
	if err := transaction.Commit().Error; err != nil {
		return domain.Transaction{}, err
	}
	transactionDomain := domain.TransactionEntityIntoDomain(transactionRaw)
	return transactionDomain, nil
}
