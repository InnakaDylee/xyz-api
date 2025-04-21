package service

import "xyz/modules/limit/domain"

type LimitCommandServiceInterface interface {
	CreateLimit(limit domain.Limit, salary int) (domain.Limit, error)
	UpdateLimit(limit domain.Limit) (domain.Limit, error)
}

type LimitQueryServiceInterface interface {
	GetLimitByConsumerId(consumerId string) ([]domain.Limit, error)
	GetLimitById(limitId string) (domain.Limit, error)
}