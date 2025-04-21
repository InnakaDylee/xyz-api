package repository

import "xyz/modules/limit/domain"

type LimitCommandRepositoryInterface interface {
	CreateLimit(limit domain.Limit) (domain.Limit, error)
	UpdateLimit(limit domain.Limit) (domain.Limit, error)
}

type LimitQueryRepositoryInterface interface {
	GetLimitByConsumerId(consumerId string) ([]domain.Limit, error)
	GetLimitById(limitId string) (domain.Limit, error)
}
