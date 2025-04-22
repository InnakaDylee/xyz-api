package service

import (
	"xyz/modules/user/domain"
	"xyz/modules/user/repository"
)

type UserQueryService struct{
	userQueryRepository repository.UserQueryRepositoryInterface
}

func NewUserQueryService(userQueryRepository repository.UserQueryRepositoryInterface) UserQueryServiceInterface {
	return &UserQueryService{
		userQueryRepository: userQueryRepository,
	}
}

func (u *UserQueryService) GetUserByID(userID int) (domain.User, error) {
	user, err := u.userQueryRepository.GetUserByID(userID)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
