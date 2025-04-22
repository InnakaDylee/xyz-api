package repository

import (
	"time"
	"xyz/modules/installment/domain"
	"xyz/modules/installment/entity"

	"gorm.io/gorm"
)

type InstallmentQueryRepository struct {
	db *gorm.DB
}

func NewInstallmentQueryRepository(db *gorm.DB) InstallmentQueryRepositoryInterface {
	return &InstallmentQueryRepository{
		db: db,
	}
}

func (r *InstallmentQueryRepository) GetInstallmentsNearDueDate(transactionId string) ([]domain.Installment, error) {
	installments := []entity.Installment{}

	now := time.Now()
	nextWeek := now.AddDate(0, 0, 7)

	if err := r.db.
		Where("transaction_id = ?", transactionId).
		Where("status = ?", "unpaid").
		Where("due_date <= ?", nextWeek).
		Order("due_date ASC").
		Find(&installments).Error; err != nil {
		return nil, err
	}

	return domain.InstallmentEntityIntoDomainList(installments), nil
}

func (r *InstallmentQueryRepository) GetInstallmentByID(id string) (domain.Installment, error) {
	var installment domain.Installment
	if err := r.db.Where("id = ?", id).
	Preload("Transaction").
	Preload("Transaction.Consumer").
	Preload("Transaction.Limit").
	First(&installment).Error; err != nil {
		return domain.Installment{}, err
	}
	return installment, nil
}

func (r *InstallmentQueryRepository) GetInstallmentsByTransactionID(transactionId string) ([]domain.Installment, error) {
	installmentEntity := []entity.Installment{}

	if err := r.db.Where("transaction_id = ?", transactionId).Find(&installmentEntity).Error; err != nil {
		return nil, err
	}
	installmentDomain := domain.InstallmentEntityIntoDomainList(installmentEntity)

	return installmentDomain, nil
}

func (r *InstallmentQueryRepository) GetInstallmentsNearDueDateWithoutIds() ([]domain.Installment, error) {
	var installments []entity.Installment

	now := time.Now()
	nextWeek := now.AddDate(0, 0, 7)

	err := r.db.
		Preload("Transaction").
		Preload("Transaction.Limit").
		Preload("Transaction.Consumer").
		Where("status = ?", "unpaid").
		Where("due_date <= ?", nextWeek).
		Order("due_date ASC").
		Find(&installments).Error

	if err != nil {
		return nil, err
	}

	return domain.InstallmentEntityIntoDomainList(installments), nil
}

