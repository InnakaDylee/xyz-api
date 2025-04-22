package service

import (
	"errors"
	"testing"

	"xyz/modules/user/domain"
	"xyz/test/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetUserByID_UserNotFound(t *testing.T) {
    mockQry := new(mocks.UserQueryRepositoryInterface)
    svc := NewUserQueryService(mockQry)

    userID := 999
    mockQry.On("GetUserByID", userID).Return(domain.User{}, errors.New("user not found"))

    user, err := svc.GetUserByID(userID)

    assert.NotNil(t, err)
    assert.Equal(t, "user not found", err.Error())
    assert.Equal(t, domain.User{}, user)
    mockQry.AssertExpectations(t)
}