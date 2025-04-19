package service

import (
	"errors"
	"fmt"
	"xyz/middlewares"
	"xyz/modules/user/domain"
	"xyz/modules/user/repository"
	"xyz/packages/bcrypt"
	"xyz/packages/util"
	"xyz/packages/validator"

	bcrypto "golang.org/x/crypto/bcrypt"
)

type userCService struct {
	userCommandRepository repository.UserCommandRepositoryInterface 
	userQueryRepository repository.UserQueryRepositoryInterface
}

func NewUserCommandService(userCommandRepository repository.UserCommandRepositoryInterface, userQueryRepository repository.UserQueryRepositoryInterface) UserCommandServiceInterface {
	return &userCService{
		userCommandRepository: userCommandRepository,
		userQueryRepository: userQueryRepository,
	}
}

func (s *userCService) Login(username, password string) (domain.User, string, error) {
	validateEmpty := validator.CheckEmpty(username, password)
	if validateEmpty != nil {
		return domain.User{}, "", validateEmpty
	}
	fmt.Println(password)
	
	user, err := s.userQueryRepository.GetUserByUsername(username)
	if err != nil {
		return domain.User{}, "", err
	}

	checkingPassword := bcrypto.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if checkingPassword != nil {
		return domain.User{}, "", errors.New("invalid password")
	}

	if !user.IsActive{
		return domain.User{}, "", errors.New("user not active")
	}

	token, err := middlewares.GenerateAuthToken(user.ID, user.Email)
	if err != nil {
		return domain.User{}, "", errors.New("failed to generate token")
	}

	return user, token, nil
}

func (s *userCService) Register(user domain.User) (domain.User, error) {
	fmt.Println(user.Password)
	validateEmpty := validator.CheckEmpty(user.Username, user.Password, user.Email)
	if validateEmpty != nil {
		return domain.User{}, validateEmpty
	}
	validateEmail := validator.CheckEmail(user.Email)
	if validateEmail != nil {
		return domain.User{}, validateEmail
	}
	validatePassword := validator.CheckPassword(user.Password)
	if validatePassword != nil {
		return domain.User{}, validatePassword
	}

	_, err := s.userQueryRepository.GetUserByUsername(user.Username)
	if err == nil {
		return domain.User{}, errors.New("email already registered")
	}

	_, err = s.userQueryRepository.GetUserByEmail(user.Email)
	if err == nil {
		return domain.User{}, errors.New("username already registered")
	}

	HashPassword, err := bcrypt.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, err
	}
	user.Password = HashPassword

	userDomain, err := s.userCommandRepository.CreateUser(user)
	if err != nil {
		return domain.User{}, err
	}

	return userDomain, nil
}

func (s *userCService) RequestActivation(email string) error {
	validateEmpty := validator.CheckEmpty(email)
	if validateEmpty != nil {
		return validateEmpty
	}
	validateEmail := validator.CheckEmail(email)
	if validateEmail != nil {
		return validateEmail
	}

	user, err := s.userQueryRepository.GetUserByEmail(email)
	if err != nil {
		return err
	}

	if user.IsActive {
		return errors.New("user already active")
	}

	token, err := middlewares.GenerateActivateToken(user.Email)
	if err != nil {
		return errors.New("failed to generate token")
	}

	go util.SendEmailNotification(user.Email, token)

	return nil
}

func (s *userCService) ActivateUser(email string) error {
	validateEmpty := validator.CheckEmpty(email)
	if validateEmpty != nil {
		return validateEmpty
	}
	validateEmail := validator.CheckEmail(email)
	if validateEmail != nil {
		return validateEmail
	}

	user, err := s.userQueryRepository.GetUserByEmail(email)
	if err != nil {
		return err
	}

	if user.IsActive {
		return errors.New("user already active")
	}

	err = s.userCommandRepository.ActivateUser(email)
	if err != nil {
		return err
	}

	return nil
}

func (s *userCService) ResetPassword(email string) (string, error) {
	return "", nil

}

func (s *userCService) ChangePassword(userID, oldPassword, newPassword string) error {
	return nil

}

