package service

import (
	"errors"
	"testing"

	"xyz/modules/limit/domain"
	"xyz/test/mocks"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
)

func TestGetLimitByConsumerId_Success(t *testing.T) {
	mockQry := new(mocks.LimitQueryRepositoryInterface)

	svc := NewLimitQueryService(mockQry)

	consumerId := "consumer-123"
	expectedLimit := []domain.Limit{
		{
			ID:              "3ec372a0-7f8a-4939-bf69-3d36fa8f7fe9",
			LimitAmount:     14000000,
			Tenor:           2,
			ConsumerID:      consumerId,
			RemainingAmount: 14000000,
		},
	}

	// Mocking the GetLimitByConsumerId method to return the expected limit
	mockQry.On("GetLimitByConsumerId", consumerId).Return(expectedLimit, nil).Once()

	// Call the GetLimitByConsumerId method
	result, err := svc.GetLimitByConsumerId(consumerId)

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, expectedLimit, result)
	mockQry.AssertExpectations(t)
}

func TestGetLimitByConsumerId_Error(t *testing.T) {
	mockQry := new(mocks.LimitQueryRepositoryInterface)

	svc := NewLimitQueryService(mockQry)

	consumerId := "consumer-123"

	// Mocking the GetLimitByConsumerId method to return an error
	mockQry.On("GetLimitByConsumerId", consumerId).Return([]domain.Limit{}, errors.New("failed to fetch limit")).Once()

	// Call the GetLimitByConsumerId method
	result, err := svc.GetLimitByConsumerId(consumerId)

	// Assertions
	assert.NotNil(t, err)
	assert.Equal(t, "failed to fetch limit", err.Error())
	assert.Equal(t, []domain.Limit{}, result)
	mockQry.AssertExpectations(t)
}
