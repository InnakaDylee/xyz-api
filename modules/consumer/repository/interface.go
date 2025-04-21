package repository

import "xyz/modules/consumer/domain"

type ConsumerCommandRepositoryInterface interface {
	CreateConsumer(consumer domain.Consumer) (domain.Consumer, error)
	UpdateConsumer(consumer domain.Consumer) (domain.Consumer, error)
}

type ConsumerQueryRepositoryInterface interface {
	GetConsumerByUserID(userID int) (domain.Consumer, error)
	GetConsumerByID(consumerID string) (domain.Consumer, error)
}