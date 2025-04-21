package repository

import (
	"fmt"
	"xyz/modules/consumer/domain"

	"gorm.io/gorm"
)

type consumerCommandRepository struct {
	db *gorm.DB
}

func NewConsumerCommandRepository(db *gorm.DB) ConsumerCommandRepositoryInterface {
	return &consumerCommandRepository{
		db: db,
	}
}

func (r *consumerCommandRepository) CreateConsumer(consumer domain.Consumer) (domain.Consumer, error) {
	transaction := r.db.Begin()
	if transaction.Error != nil {
		return domain.Consumer{}, transaction.Error
	}
	consumerRaw := domain.ConsumerDomainIntoEntity(consumer)
	if err := transaction.Create(&consumerRaw).Error; err != nil {

		transaction.Rollback()
		return domain.Consumer{}, err
	}

	consumerDomain := domain.ConsumerEntityIntoDomain(consumerRaw)
	if err := transaction.Commit().Error; err != nil {
		return domain.Consumer{}, err
	}

	return consumerDomain, nil
}

func (r *consumerCommandRepository) UpdateConsumer(consumer domain.Consumer) (domain.Consumer, error) {
	transaction := r.db.Begin()
	if transaction.Error != nil {
		return domain.Consumer{}, transaction.Error
	}
	consumerRaw := domain.ConsumerDomainIntoEntity(consumer)
	if err := transaction.Save(&consumerRaw).Error; err != nil {
		fmt.Println("error save consumer", err)
		transaction.Rollback()
		return domain.Consumer{}, err
	}

	if err := transaction.Commit().Error; err != nil {
		return domain.Consumer{}, err
	}

	consumerDomain := domain.ConsumerEntityIntoDomain(consumerRaw)
	if err := transaction.Commit().Error; err != nil {
		return domain.Consumer{}, err
	}

	return consumerDomain, nil
}