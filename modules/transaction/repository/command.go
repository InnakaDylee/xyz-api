package repository

import (
	"fmt"
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
	fmt.Println("masuk create transaction")
	transaction := r.db.Begin()
	if transaction.Error != nil {
		return domain.Transaction{}, transaction.Error
	}
	fmt.Println("masuk create transaction 2")

	transactionRaw := domain.TransactionDomainIntoEntity(transactionInput)
	fmt.Println("masuk create transaction 3")
	if err := transaction.Debug().Create(&transactionRaw).Error; err != nil {
		fmt.Println("error create transaction", err)
		transaction.Rollback()
		return domain.Transaction{}, err
	}
	fmt.Println("masuk create transaction 4")
	if err := transaction.Commit().Error; err != nil {
		return domain.Transaction{}, err
	}
	fmt.Println("masuk create transaction 5")
	transactionDomain := domain.TransactionEntityIntoDomain(transactionRaw)
	return transactionDomain, nil
}
