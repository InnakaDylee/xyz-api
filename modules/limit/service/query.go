package service

import (
	"xyz/modules/limit/repository"
	"xyz/modules/limit/domain"
)

type LimitQueryService struct {
	limitQueryRepository repository.LimitQueryRepositoryInterface
}

func NewLimitQueryService(limitQueryRepository repository.LimitQueryRepositoryInterface) LimitQueryServiceInterface {
	return &LimitQueryService{
		limitQueryRepository: limitQueryRepository,
	}
}

func (s *LimitQueryService) GetLimitByConsumerId(consumerId string) ([]domain.Limit, error) {
	limit, err := s.limitQueryRepository.GetLimitByConsumerId(consumerId)
	if err != nil {
		return []domain.Limit{}, err
	}
	return limit, nil
}

func (s *LimitQueryService) GetLimitById(limitId string) (domain.Limit, error) {
	limit, err := s.limitQueryRepository.GetLimitById(limitId)
	if err != nil {
		return domain.Limit{}, err
	}
	return limit, nil
}