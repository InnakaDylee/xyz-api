package repository

import (
	"errors"
	"xyz/modules/consumer/domain"
	"xyz/modules/consumer/entity"

	"gorm.io/gorm"
)

type consumerQueryRepository struct {
	db *gorm.DB
}

func NewConsumerQueryRepository(db *gorm.DB) ConsumerQueryRepositoryInterface {
	return &consumerQueryRepository{
		db: db,
	}
}

func (r *consumerQueryRepository) GetConsumerByID(consumerID string) (domain.Consumer, error) {
	consumer := entity.Consumer{}
	if err := r.db.Where("id = ?", consumerID).First(&consumer).Error; err != nil {
		return domain.Consumer{}, err
	}
	if consumer == (entity.Consumer{}) {
		return domain.Consumer{}, errors.New("consumer not found")
	}

	consumerDomain := domain.ConsumerEntityIntoDomain(consumer)

	return consumerDomain, nil
}

func (r *consumerQueryRepository) GetConsumerByUserID(userID int) (domain.Consumer, error) {
	consumer := entity.Consumer{}
	if err := r.db.Where("user_id = ?", userID).First(&consumer).Error; err != nil {
		return domain.Consumer{}, err
	}
	if consumer == (entity.Consumer{}) {
		return domain.Consumer{}, errors.New("consumer not found")
	}

	consumerDomain := domain.ConsumerEntityIntoDomain(consumer)

	return consumerDomain, nil
}