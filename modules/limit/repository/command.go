package repository

import (
	"xyz/modules/limit/domain"

	"gorm.io/gorm"
)

type LimitCommandRepository struct {
	db *gorm.DB
}

func NewLimitCommandRepository(db *gorm.DB) *LimitCommandRepository {
	return &LimitCommandRepository{
		db: db,
	}
}

func (r *LimitCommandRepository) CreateLimit(limit domain.Limit) (domain.Limit, error) {
	transaction := r.db.Begin()
	if transaction.Error != nil {
		return domain.Limit{}, transaction.Error
	}
	limitRaw := domain.LimitDomainIntoEntity(limit)
	if err := transaction.Create(&limitRaw).Error; err != nil {
		transaction.Rollback()
		return domain.Limit{}, err
	}

	limitDomain := domain.LimitEntityIntoDomain(limitRaw)
	if err := transaction.Commit().Error; err != nil {
		return domain.Limit{}, err
	}

	return limitDomain, nil
}

func (r *LimitCommandRepository) UpdateLimit(limit domain.Limit) (domain.Limit, error) {
	transaction := r.db.Begin()
	if transaction.Error != nil {
		return domain.Limit{}, transaction.Error
	}

	limitRaw := domain.LimitDomainIntoEntityWithLimitID(limit)
	if err := transaction.Save(&limitRaw).Error; err != nil {
		transaction.Rollback()
		return domain.Limit{}, err
	}

	if err := transaction.Commit().Error; err != nil {
		return domain.Limit{}, err
	}

	limitDomain := domain.LimitEntityIntoDomain(limitRaw)

	return limitDomain, nil
}
