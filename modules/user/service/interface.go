package service

import "xyz/modules/user/domain"

type UserCommandServiceInterface interface {
	Login(username, password string) (domain.User, string, error)
	Register(domain.User) (domain.User,  error)
	RequestActivation(email string) error
	ActivateUser(userID string) error
	ResetPassword(email string) (string, error)
	ChangePassword(userID, oldPassword, newPassword string) error
}

type UserQueryServiceInterface interface {
	GetUserByID(userID string) (domain.User, error)
}
