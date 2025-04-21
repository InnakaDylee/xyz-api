package repository

import (
	"xyz/modules/limit/entity"
	"xyz/modules/limit/domain"

	"gorm.io/gorm"
)

type LimitQueryRepository struct {
	db *gorm.DB
}

func NewLimitQueryRepository(db *gorm.DB) *LimitQueryRepository {
	return &LimitQueryRepository{
		db: db,
	}
}

func (r *LimitQueryRepository) GetLimitById(limitId string) (domain.Limit, error) {
	var limit entity.Limit

	if err := r.db.Where("id = ?", limitId).Preload("Consumer").First(&limit).Error; err != nil {
		return domain.Limit{}, err
	}

	limitDomain := domain.LimitEntityIntoDomainWithConsumer(limit)

	return limitDomain, nil
}

func (r *LimitQueryRepository) GetLimitByConsumerId(consumerId string) ([]domain.Limit, error) {
	var limit []entity.Limit
	if err := r.db.Debug().Where("consumer_id = ?", consumerId).Find(&limit).Error; err != nil {
		return []domain.Limit{}, err
	}
	limitDomain := domain.ManyLimitEntityIntoDomain(limit)

	return limitDomain, nil
}