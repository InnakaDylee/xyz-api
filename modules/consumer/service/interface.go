package service

import (
	"mime/multipart"
	"xyz/modules/consumer/domain"
)

type ConsumerCommandServiceInterface interface {
	CreateConsumer(consumer domain.Consumer, photoKTP, photoSelfie *multipart.FileHeader) (domain.Consumer, error)
	UpdateConsumer(consumer domain.Consumer, photoKTP, photoSelfie *multipart.FileHeader) (domain.Consumer, error)
}

type ConsumerQueryServiceInterface interface {
	GetConsumerByUserID(userID int) (domain.Consumer, error)
	GetConsumerByID(consumerID string) (domain.Consumer, error)
}