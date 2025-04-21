package service

import (
	"xyz/modules/consumer/domain"
	"xyz/modules/consumer/repository"
)

type ConsumerQueryService struct {
	ConsumerQueryRepository repository.ConsumerQueryRepositoryInterface
}

func NewConsumerQueryService(consumerQueryRepository repository.ConsumerQueryRepositoryInterface) ConsumerQueryServiceInterface {
	return &ConsumerQueryService{
		ConsumerQueryRepository: consumerQueryRepository,
	}
}

func (c *ConsumerQueryService) GetConsumerByUserID(userID int) (domain.Consumer, error) {
	consumer, err := c.ConsumerQueryRepository.GetConsumerByUserID(userID)
	if err != nil {
		return domain.Consumer{}, err
	}
	return consumer, nil
}

func (c *ConsumerQueryService) GetConsumerByID(consumerID string) (domain.Consumer, error) {
	consumer, err := c.ConsumerQueryRepository.GetConsumerByID(consumerID)
	if err != nil {
		return domain.Consumer{}, err
	}
	return consumer, nil
}