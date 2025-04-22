package repository

import (
	"errors"
	"xyz/modules/user/domain"
	"xyz/modules/user/entity"

	"gorm.io/gorm"
)

type userQueryRepository struct{
	db *gorm.DB
}

func NewUserQueryRepository(db *gorm.DB) UserQueryRepositoryInterface {
	return &userQueryRepository{
		db: db,
	}
}

func (r *userQueryRepository) GetUserByID(userID int) (domain.User, error) {
	user := entity.User{}
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	if user == (entity.User{}) {
		return domain.User{}, errors.New("user not found")
	}

	userDomain := domain.UserEntityIntoDomain(user)
	
	return userDomain, nil
}

func (r *userQueryRepository) GetUserByUsername(username string) (domain.User, error) {
	user := entity.User{}
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	if user == (entity.User{}) {
		return domain.User{}, errors.New("user not found")
	}

	userDomain := domain.UserEntityIntoDomain(user)
	
	return userDomain, nil
}

func (r *userQueryRepository) GetUserByEmail(email string) (domain.User, error) {
	user := entity.User{}
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return domain.User{}, errors.New("user not found")
	}
	if user == (entity.User{}) {
		return domain.User{}, errors.New("user not found")
	}

	userDomain := domain.UserEntityIntoDomain(user)
	
	return userDomain, nil
}