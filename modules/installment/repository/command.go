package repository

import (
	"xyz/modules/installment/domain"

	"gorm.io/gorm"
)

type InstallmentCommandRepository struct {
	db *gorm.DB
}

func NewInstallmentCommandRepository(db *gorm.DB) InstallmentCommandRepositoryInterface {
	return &InstallmentCommandRepository{
		db: db,
	}
}

func (r *InstallmentCommandRepository) CreateInstallment(installment domain.Installment) (domain.Installment, error) {
	installmentEntity := domain.InstallmentDomainIntoEntity(installment)
	transaction := r.db.Begin()
	if transaction.Error != nil {
		return domain.Installment{}, transaction.Error
	}

	if err := transaction.Create(&installmentEntity).Error; err != nil {
		transaction.Rollback()
		return domain.Installment{}, err
	}

	if err := transaction.Commit().Error; err != nil {
		return domain.Installment{}, err
	}

	installment = domain.InstallmentEntityIntoDomain(installmentEntity)

	return installment, nil
}

func (r *InstallmentCommandRepository) UpdateInstallment(installment domain.Installment) (domain.Installment, error) {
	transaction := r.db.Begin()
	if transaction.Error != nil {
		return domain.Installment{}, transaction.Error
	}

	if err := transaction.Save(&installment).Error; err != nil {
		transaction.Rollback()
		return domain.Installment{}, err
	}

	if err := transaction.Commit().Error; err != nil {
		return domain.Installment{}, err
	}

	return installment, nil
}