package domain

import (
	"time"
	"xyz/modules/consumer/domain"
	"xyz/modules/limit/entity"
)

type Limit struct {
	ID          string `json:"id"`
	LimitAmount int    `json:"limit_amount"`
	Tenor 	 	int    `json:"tenor"`
	ConsumerID  string `json:"consumer_id"`
	Consumer	domain.Consumer
	RemainingAmount int `json:"remaining_amount"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewLimit(consumer domain.Consumer) Limit {
	return Limit{
		ConsumerID:             consumer.ID,
	}
}

func LimitDomainIntoEntity(limit Limit) entity.Limit {
	return entity.Limit{
		LimitAmount:    limit.LimitAmount,
		Tenor:          limit.Tenor,
		ConsumerID:     limit.ConsumerID,
		RemainingAmount: limit.RemainingAmount,
		CreatedAt:      limit.CreatedAt,
		UpdatedAt:      limit.UpdatedAt,
	}
}

func LimitDomainIntoEntityWithLimitID(limit Limit) entity.Limit {
	return entity.Limit{
		ID:             limit.ID,
		LimitAmount:    limit.LimitAmount,
		Tenor:          limit.Tenor,
		ConsumerID:     limit.ConsumerID,
		RemainingAmount: limit.RemainingAmount,
		CreatedAt:      limit.CreatedAt,
		UpdatedAt:      limit.UpdatedAt,
	}
}

func LimitEntityIntoDomain(limit entity.Limit) Limit {
	return Limit{
		ID:             limit.ID,
		LimitAmount:    limit.LimitAmount,
		Tenor:          limit.Tenor,
		ConsumerID:     limit.ConsumerID,
		RemainingAmount: limit.RemainingAmount,
		CreatedAt:      limit.CreatedAt,
		UpdatedAt:      limit.UpdatedAt,
	}
}

func LimitEntityIntoDomainWithConsumer(limit entity.Limit) Limit {
	return Limit{
		ID:             limit.ID,
		LimitAmount:    limit.LimitAmount,
		Tenor:          limit.Tenor,
		ConsumerID:     limit.ConsumerID,
		Consumer:       domain.ConsumerEntityIntoDomain(limit.Consumer),
		RemainingAmount: limit.RemainingAmount,
		CreatedAt:      limit.CreatedAt,
		UpdatedAt:      limit.UpdatedAt,
	}
}

func ManyLimitEntityIntoDomain(limit []entity.Limit) []Limit {
	limitDomain := []Limit{}
	for _, limit := range limit {
		limitDomain = append(limitDomain, Limit{
			ID:             limit.ID,
			LimitAmount:    limit.LimitAmount,
			Tenor:          limit.Tenor,
			ConsumerID:     limit.ConsumerID,
			RemainingAmount: limit.RemainingAmount,
			CreatedAt:      limit.CreatedAt,
			UpdatedAt:      limit.UpdatedAt,
		})
	}
	return limitDomain
}

