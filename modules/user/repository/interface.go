package repository

import "xyz/modules/user/domain"

type UserCommandRepositoryInterface interface {
	CreateUser(user domain.User) (domain.User, error)
	UpdateUser(user domain.User) error
	ActivateUser(userID string) error
}

type UserQueryRepositoryInterface interface {
	GetUserByID(userID string) (domain.User, error)
	GetUserByUsername(username string) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
}