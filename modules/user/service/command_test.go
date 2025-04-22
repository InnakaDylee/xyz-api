package service

import (
	"errors"
	"testing"

	"xyz/modules/user/domain"
	"xyz/packages/bcrypt"
	"xyz/test/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


func TestLogin_Success(t *testing.T) {
    mockCmd := new(mocks.UserCommandRepositoryInterface)
    mockQry := new(mocks.UserQueryRepositoryInterface)
    svc := NewUserCommandService(mockCmd, mockQry)

    hashed, _ := bcrypt.HashPassword("Password123!")

    mockQry.On("GetUserByUsername", "johndoe").Return(domain.User{
        ID:       1,
        Username: "johndoe",
        Email:    "john@mail.com",
        Password: hashed,
        IsActive: true,
    }, nil)

    _, _, err := svc.Login("johndoe", "Password123!")

    assert.Nil(t, err)
}

func TestLogin_InvalidPassword(t *testing.T) {
	mockCmd := new(mocks.UserCommandRepositoryInterface)
	mockQry := new(mocks.UserQueryRepositoryInterface)
	svc := NewUserCommandService(mockCmd, mockQry) 

	hashed, _ := bcrypt.HashPassword("Password123!")

	mockQry.On("GetUserByUsername", "johndoe").Return(domain.User{
		ID:       1,
		Username: "johndoe",
		Email:    "john@gmail.com",
		Password: hashed,
		IsActive: true,
	}, nil)

	_, _, err := svc.Login("johndoe", "wrongpassword")

	assert.NotNil(t, err)
	assert.Equal(t, "invalid password", err.Error())
	mockQry.AssertExpectations(t)
}

func TestLogin_UserNotActive(t *testing.T) {
	mockCmd := new(mocks.UserCommandRepositoryInterface)
	mockQry := new(mocks.UserQueryRepositoryInterface)
	svc := NewUserCommandService(mockCmd, mockQry)

	hashed, _ := bcrypt.HashPassword("Password123!")

	mockQry.On("GetUserByUsername", "johndoe").Return(domain.User{
		ID:       1,
		Username: "johndoe",
		Email:    "john@gmail.com",
		Password: hashed,
		IsActive: false,
	}, nil)
	_, _, err := svc.Login("johndoe", "Password123!")
	assert.NotNil(t, err)
	assert.Equal(t, "user not active", err.Error())
	mockQry.AssertExpectations(t)
}

func TestLogin_EmptyUsername(t *testing.T) {
	mockCmd := new(mocks.UserCommandRepositoryInterface)
	mockQry := new(mocks.UserQueryRepositoryInterface)
	svc := NewUserCommandService(mockCmd, mockQry) 

	_, _, err := svc.Login("", "Password123!")

	assert.NotNil(t, err)
	assert.Equal(t, "field cannot be empty", err.Error())
}

func TestLogin_EmptyPassword(t *testing.T) {
	mockCmd := new(mocks.UserCommandRepositoryInterface)
	mockQry := new(mocks.UserQueryRepositoryInterface)
	svc := NewUserCommandService(mockCmd, mockQry) 

	_, _, err := svc.Login("johndoe", "")

	assert.NotNil(t, err)
	assert.Equal(t, "field cannot be empty", err.Error())
}

func TestRegister_Success(t *testing.T) {
	mockCmd := new(mocks.UserCommandRepositoryInterface)
	mockQry := new(mocks.UserQueryRepositoryInterface)
	svc := NewUserCommandService(mockCmd, mockQry)

	userInput := domain.User{
		Username: "johndoe",
		Email:    "john@example.com",
		Password: "Password123!",
	}


	mockQry.On("GetUserByUsername", userInput.Username).Return(domain.User{}, errors.New("not found"))
	mockQry.On("GetUserByEmail", userInput.Email).Return(domain.User{}, errors.New("not found"))
	mockCmd.On("CreateUser", mock.AnythingOfType("domain.User")).Return(userInput, nil)


	createdUser, err := svc.Register(userInput)


	assert.Nil(t, err)
	assert.Equal(t, userInput.Username, createdUser.Username)
	assert.Equal(t, userInput.Email, createdUser.Email)

	mockQry.AssertExpectations(t)
	mockCmd.AssertExpectations(t)
}

func TestRegister_UsernameAlreadyExists(t *testing.T) {
	mockCmd := new(mocks.UserCommandRepositoryInterface)
	mockQry := new(mocks.UserQueryRepositoryInterface)
	svc := NewUserCommandService(mockCmd, mockQry)

	userInput := domain.User{
		Username: "johndoe",
		Email:    "john@example.com",
		Password: "Password123!",
	}


	mockQry.On("GetUserByUsername", userInput.Username).Return(userInput, nil)


	_, err := svc.Register(userInput)


	assert.NotNil(t, err)
	assert.Equal(t, "username already registered", err.Error())

	mockQry.AssertExpectations(t)
}

func TestRegister_EmailAlreadyExists(t *testing.T) {
	mockCmd := new(mocks.UserCommandRepositoryInterface)
	mockQry := new(mocks.UserQueryRepositoryInterface)
	svc := NewUserCommandService(mockCmd, mockQry)

	userInput := domain.User{
		Username: "johndoe",
		Email:    "john@example.com",
		Password: "Password123!",
	}


	mockQry.On("GetUserByUsername", userInput.Username).Return(domain.User{}, errors.New("not found"))
	mockQry.On("GetUserByEmail", userInput.Email).Return(userInput, nil)

	_, err := svc.Register(userInput)

	assert.NotNil(t, err)
	assert.Equal(t, "email already registered", err.Error())
	mockQry.AssertExpectations(t)
	mockCmd.AssertExpectations(t)
}

func TestRegister_EmptyUsername(t *testing.T) {
	mockCmd := new(mocks.UserCommandRepositoryInterface)
	mockQry := new(mocks.UserQueryRepositoryInterface)
	svc := NewUserCommandService(mockCmd, mockQry)

	userInput := domain.User{
		Username: "",
		Email:    "john@gmail.com",
		Password: "Password123!",
	}

	_, err := svc.Register(userInput)

	assert.NotNil(t, err)
	assert.Equal(t, "field cannot be empty", err.Error())
	mockQry.AssertExpectations(t)
	mockCmd.AssertExpectations(t)
}

func TestRequestActivation_UserAlreadyActive(t *testing.T) {
    mockCmd := new(mocks.UserCommandRepositoryInterface)
    mockQry := new(mocks.UserQueryRepositoryInterface)
    svc := NewUserCommandService(mockCmd, mockQry)

    email := "john@mail.com"
    user := domain.User{
        Email:    email,
        IsActive: true,
    }

    mockQry.On("GetUserByEmail", email).Return(user, nil)

    err := svc.RequestActivation(email)

    assert.NotNil(t, err)
    assert.Equal(t, "user already active", err.Error())
    mockQry.AssertExpectations(t)
}

func TestRequestActivation_EmptyEmail(t *testing.T) {
    mockCmd := new(mocks.UserCommandRepositoryInterface)
    mockQry := new(mocks.UserQueryRepositoryInterface)
    svc := NewUserCommandService(mockCmd, mockQry)

    email := ""

    err := svc.RequestActivation(email)

    assert.NotNil(t, err)
    assert.Equal(t, "field cannot be empty", err.Error())
}

func TestRequestActivation_InvalidEmailFormat(t *testing.T) {
    mockCmd := new(mocks.UserCommandRepositoryInterface)
    mockQry := new(mocks.UserQueryRepositoryInterface)
    svc := NewUserCommandService(mockCmd, mockQry)

    email := "invalid-email"

    err := svc.RequestActivation(email)

    assert.NotNil(t, err)
    assert.Equal(t, "invalid email format", err.Error())
}

func TestRequestActivation_UserNotFound(t *testing.T) {
    mockCmd := new(mocks.UserCommandRepositoryInterface)
    mockQry := new(mocks.UserQueryRepositoryInterface)
    svc := NewUserCommandService(mockCmd, mockQry)

    email := "notfound@example.com"
    mockQry.On("GetUserByEmail", email).Return(domain.User{}, errors.New("user not found"))

    err := svc.RequestActivation(email)

    assert.NotNil(t, err)
    assert.Equal(t, "user not found", err.Error())
    mockQry.AssertExpectations(t)
}

func TestActivateUser_EmptyEmail(t *testing.T) {
    mockCmd := new(mocks.UserCommandRepositoryInterface)
    mockQry := new(mocks.UserQueryRepositoryInterface)
    svc := NewUserCommandService(mockCmd, mockQry)

    err := svc.ActivateUser("")

    assert.NotNil(t, err)
    assert.Equal(t, "field cannot be empty", err.Error())
}

func TestActivateUser_InvalidEmailFormat(t *testing.T) {
    mockCmd := new(mocks.UserCommandRepositoryInterface)
    mockQry := new(mocks.UserQueryRepositoryInterface)
    svc := NewUserCommandService(mockCmd, mockQry)

    err := svc.ActivateUser("invalid-email")

    assert.NotNil(t, err)
    assert.Equal(t, "invalid email format", err.Error())
}

func TestActivateUser_UserNotFound(t *testing.T) {
    mockCmd := new(mocks.UserCommandRepositoryInterface)
    mockQry := new(mocks.UserQueryRepositoryInterface)
    svc := NewUserCommandService(mockCmd, mockQry)

    email := "notfound@example.com"
    mockQry.On("GetUserByEmail", email).Return(domain.User{}, errors.New("user not found"))

    err := svc.ActivateUser(email)
    assert.NotNil(t, err)
    assert.Equal(t, "user not found", err.Error())
    mockQry.AssertExpectations(t)
}

func TestActivateUser_AlreadyActive(t *testing.T) {
    mockCmd := new(mocks.UserCommandRepositoryInterface)
    mockQry := new(mocks.UserQueryRepositoryInterface)
    svc := NewUserCommandService(mockCmd, mockQry)

    email := "active@example.com"
    mockQry.On("GetUserByEmail", email).Return(domain.User{
        Email:    email,
        IsActive: true,
    }, nil)

    err := svc.ActivateUser(email)
    assert.NotNil(t, err)
    assert.Equal(t, "user already active", err.Error())
    mockQry.AssertExpectations(t)
}
