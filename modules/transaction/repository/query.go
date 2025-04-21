package repository

import (
	"xyz/modules/transaction/domain"

	"gorm.io/gorm"
)

type TransactionQueryRepository struct {
	db *gorm.DB
}

func NewTransactionQueryRepository(db *gorm.DB) TransactionQueryRepositoryInterface {
	return &TransactionQueryRepository{
		db: db,
	}
}

func (r *TransactionQueryRepository) GetTransactionByID(id string) (domain.Transaction, error) {
	var transaction domain.Transaction
	if err := r.db.Where("id = ?", id).Preload("Limit").Preload("Consumer").First(&transaction).Error; err != nil {
		return domain.Transaction{}, err
	}
	return transaction, nil
}

func (r *TransactionQueryRepository) GetTransactionByConsumerID(consumerID string) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	if err := r.db.Where("consumer_id = ?", consumerID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}