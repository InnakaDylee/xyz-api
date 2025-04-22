package service

import (
	"testing"
	"time"

	"xyz/modules/consumer/domain"
	"xyz/test/mocks"

	"github.com/stretchr/testify/assert"
)

func timeMustParse(value string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", value)
	if err != nil {
		panic(err)
	}
	return t
}

func TestGetConsumerByUserID_Success(t *testing.T) {
	mockQry := new(mocks.ConsumerQueryRepositoryInterface)
	svc := NewConsumerQueryService(mockQry)
	
	userID := 5
	consumerExp := domain.Consumer{
		ID:            "3ec372a0-7f8a-4939-bf69-3d36fa8f7fe9",
		User_ID:        userID,
		NIK: 		 "1234567890123456",
		Full_Name:     "jhon doe",
		Legal_Name:    "jhon doe",
		Place_Of_Birth: "kuala tungkal",
		Date_Of_Birth:  "01-01-2025",
		Salary:        10000000,
		Photo_KTP:     "storages/images/8Bpibdx6lT_ktpbre.jpeg",
		Photo_Selfie:  "storages/images/urDFgphxfw_selfie.jpeg",
		CreatedAt:    timeMustParse("2025-04-20 15:35:36"),
		UpdatedAt:    timeMustParse("2025-04-20 15:35:36"),
	}

	mockQry.On("GetConsumerByUserID", userID).Return(consumerExp, nil)

	consumer, err := svc.GetConsumerByUserID(userID)
	assert.Nil(t, err)
	assert.Equal(t, consumerExp, consumer)
	mockQry.AssertExpectations(t)
}

func TestGetConsumerByUserID_UserNotFound(t *testing.T) {
	mockQry := new(mocks.ConsumerQueryRepositoryInterface)
	svc := NewConsumerQueryService(mockQry)

	userID := 999
	mockQry.On("GetConsumerByUserID", userID).Return(domain.Consumer{}, nil)

	consumer, err := svc.GetConsumerByUserID(userID)
	assert.Nil(t, err)
	assert.NotNil(t, consumer)
}

func TestGetConsumerByID_Success(t *testing.T) {
	mockQry := new(mocks.ConsumerQueryRepositoryInterface)
	svc := NewConsumerQueryService(mockQry)

	consumerID := "3ec372a0-7f8a-4939-bf69-3d36fa8f7fe9"
	consumerExp := domain.Consumer{
		ID:            consumerID,
		User_ID:        5,
		NIK: 		 "1234567890123456",
		Full_Name:     "jhon doe",
		Legal_Name:    "jhon doe",
		Place_Of_Birth: "kuala tungkal",
		Date_Of_Birth:  "01-01-2025",
		Salary:        10000000,
		Photo_KTP:     "storages/images/8Bpibdx6lT_ktpbre.jpeg",
		Photo_Selfie:  "storages/images/urDFgphxfw_selfie.jpeg",
		CreatedAt:    timeMustParse("2025-04-20 15:35:36"),
		UpdatedAt:    timeMustParse("2025-04-20 15:35:36"),
	}

	mockQry.On("GetConsumerByID", consumerID).Return(consumerExp, nil)

	consumer, err := svc.GetConsumerByID(consumerID)
	assert.Nil(t, err)
	assert.Equal(t, consumerExp, consumer)
	mockQry.AssertExpectations(t)
}

func TestGetConsumerByID_UserNotFound(t *testing.T) {
	mockQry := new(mocks.ConsumerQueryRepositoryInterface)
	svc := NewConsumerQueryService(mockQry)

	consumerID := "999"
	mockQry.On("GetConsumerByID", consumerID).Return(domain.Consumer{}, nil)

	consumer, err := svc.GetConsumerByID(consumerID)
	assert.Nil(t, err)
	assert.NotNil(t, consumer)
}