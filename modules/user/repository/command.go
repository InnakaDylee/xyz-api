package repository

import (
	"xyz/modules/user/domain"

	"gorm.io/gorm"
)

type userCommandRepository struct {
	db *gorm.DB
}

func NewUserCommandRepository(db *gorm.DB) UserCommandRepositoryInterface {
	return &userCommandRepository{
		db: db,
	}
}

func (r *userCommandRepository) CreateUser(user domain.User) (domain.User, error) {
	transaction := r.db.Begin()
	if transaction.Error != nil {
		return domain.User{}, transaction.Error
	}
	userRaw := domain.UserDomainIntoEntity(user)
	if err := transaction.Create(&userRaw).Error; err != nil {
		transaction.Rollback()
		return domain.User{}, err
	}
	
	userDomain := domain.UserEntityIntoDomain(userRaw)
	if err := transaction.Commit().Error; err != nil {
		return domain.User{}, err
	}

	return userDomain, nil
}

func (r *userCommandRepository) UpdateUser(user domain.User) error {
	transaction := r.db.Begin()
	if transaction.Error != nil {
		return transaction.Error
	}
	userRaw := domain.UserDomainIntoEntity(user)
	if err := transaction.Save(&userRaw).Error; err != nil {
		transaction.Rollback()
		return err
	}
	
	if err := transaction.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r *userCommandRepository) ActivateUser(email string) error {
	transaction := r.db.Begin()
	if transaction.Error != nil {
		return transaction.Error
	}
	
	if err := transaction.Model(&domain.User{}).Where("email = ?", email).Update("is_active", true).Error; err != nil {
		transaction.Rollback()
		return err
	}
	
	if err := transaction.Commit().Error; err != nil {
		return err
	}

	return nil
}