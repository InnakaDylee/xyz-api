package service

import (
	"errors"
	"testing"

	"xyz/modules/limit/domain"
	"xyz/test/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateLimit_Success(t *testing.T) {
	mockCmd := new(mocks.LimitCommandRepositoryInterface)
	mockQry := new(mocks.LimitQueryRepositoryInterface)
	svc := NewLimitCommandService(mockCmd, mockQry)

	salary := 10_000_000
	consumerID := "consumer-123"

	// Mock CreateLimit akan dipanggil 4 kali untuk tenor 1, 2, 3, dan 6
	mockCmd.On("CreateLimit", mock.MatchedBy(func(l domain.Limit) bool {
		return l.ConsumerID == consumerID && l.LimitAmount > 0 && l.Tenor > 0 && l.RemainingAmount == l.LimitAmount
	})).Return(domain.Limit{}, nil).Times(4)

	// Dummy input limit
	inputLimit := domain.Limit{
		ID:         "3ec372a0-7f8a-4939-bf69-3d36fa8f7fe9",
		ConsumerID: consumerID,
	}

	// Call the actual service
	result, err := svc.CreateLimit(inputLimit, salary)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, inputLimit.ConsumerID, result.ConsumerID)
	assert.NotZero(t, result.LimitAmount)
	assert.NotZero(t, result.Tenor)
	assert.Equal(t, result.LimitAmount, result.RemainingAmount)

	// Verify mock expectations
	mockCmd.AssertExpectations(t)
}

func TestCreateLimit_FailOnRepository(t *testing.T) {
	mockCmd := new(mocks.LimitCommandRepositoryInterface)
	mockQry := new(mocks.LimitQueryRepositoryInterface)
	svc := NewLimitCommandService(mockCmd, mockQry)

	salary := 10_000_000
	consumerID := "consumer-123"

	mockCmd.On("CreateLimit", mock.Anything).Return(domain.Limit{}, errors.New("db error")).Once()

	inputLimit := domain.Limit{
		ID:         "some-id",
		ConsumerID: consumerID,
	}

	result, err := svc.CreateLimit(inputLimit, salary)

	assert.Error(t, err)
	assert.Equal(t, "db error", err.Error())
	assert.Empty(t, result.ID)

	mockCmd.AssertExpectations(t)
}

func TestUpdateLimit_Success(t *testing.T) {
	mockCmd := new(mocks.LimitCommandRepositoryInterface)
	mockQry := new(mocks.LimitQueryRepositoryInterface)

	svc := NewLimitCommandService(mockCmd, mockQry)

	limit := domain.Limit{
		ID:              "3ec372a0-7f8a-4939-bf69-3d36fa8f7fe9",
		LimitAmount:     14000000,
		Tenor:           2,
		ConsumerID:      "consumer-123",
		RemainingAmount: 14000000,
	}

	// Mocking the UpdateLimit function to return the same limit object with no error
	mockCmd.On("UpdateLimit", mock.Anything).Return(limit, nil).Once()

	// Call the UpdateLimit method
	result, err := svc.UpdateLimit(limit)

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, limit, result)
	mockCmd.AssertExpectations(t)
}

func TestUpdateLimit_Error(t *testing.T) {
	mockCmd := new(mocks.LimitCommandRepositoryInterface)
	mockQry := new(mocks.LimitQueryRepositoryInterface)

	svc := NewLimitCommandService(mockCmd, mockQry)

	limit := domain.Limit{
		ID:              "3ec372a0-7f8a-4939-bf69-3d36fa8f7fe9",
		LimitAmount:     14000000,
		Tenor:           2,
		ConsumerID:      "consumer-123",
		RemainingAmount: 14000000,
	}

	// Simulate error returned by UpdateLimit
	mockCmd.On("UpdateLimit", mock.Anything).Return(domain.Limit{}, errors.New("failed to update")).Once()

	// Call the UpdateLimit method
	result, err := svc.UpdateLimit(limit)

	// Assertions
	assert.NotNil(t, err)
	assert.Equal(t, "failed to update", err.Error())
	assert.Equal(t, domain.Limit{}, result)
	mockCmd.AssertExpectations(t)
}
