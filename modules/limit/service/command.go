package service

import (
	"errors"
	"xyz/modules/limit/domain"
	"xyz/modules/limit/repository"
)

type LimitCommandService struct {
	LimitCommandRepositoryInterface repository.LimitCommandRepositoryInterface
	LimitQueryRepositoryInterface    repository.LimitQueryRepositoryInterface
}

func NewLimitCommandService(
	limitCommandRepositoryInterface repository.LimitCommandRepositoryInterface,
	limitQueryRepositoryInterface repository.LimitQueryRepositoryInterface,
) LimitCommandServiceInterface {
	return &LimitCommandService{
		LimitCommandRepositoryInterface: limitCommandRepositoryInterface,
		LimitQueryRepositoryInterface:    limitQueryRepositoryInterface,
	}
}

func (s *LimitCommandService) CreateLimit(limit domain.Limit, salary int) (domain.Limit, error) {
	tenorMultiply := map[int]float64{
		1: 0.8,
		2: 1.4,  
		3: 1.8,  
		6: 2.4,
	}
	var tenorMultiplyValue float64
	var ok bool 
	var err error

	for i := 1; i <= 6; i++ {
		tenorMultiplyValue, ok = tenorMultiply[i]
		if ok {
			limitAmount := int(float64(salary) * tenorMultiplyValue)
			if limitAmount < 0 {
				return domain.Limit{}, errors.New("limit amount cannot be negative")
			}
			limit.LimitAmount = limitAmount
			limit.Tenor = i
			limit.RemainingAmount = limitAmount

			_, err = s.LimitCommandRepositoryInterface.CreateLimit(limit)
			if err != nil {
				return domain.Limit{}, err
			}
		}
	}

	return limit, nil
}

func (s *LimitCommandService) UpdateLimit(limit domain.Limit) (domain.Limit, error) {
	limit, err := s.LimitCommandRepositoryInterface.UpdateLimit(limit)
	if err != nil {
		return domain.Limit{}, err
	}
	return limit, nil
}