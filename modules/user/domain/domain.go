package domain

import (
	"time"
	"xyz/modules/user/entity"
)

type User struct{
	ID int
	Email string
	Username string
	Password string
	IsActive bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func UserDomainIntoEntity(user User) (entity.User) {
	return entity.User{
		ID: user.ID,
		Email: user.Email,
		Username: user.Username,
		Password: user.Password,
		IsActive: user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserEntityIntoDomain(user entity.User) (User) {
	return User{
		ID: user.ID,
		Email: user.Email,
		Username: user.Username,
		Password: user.Password,
		IsActive: user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}