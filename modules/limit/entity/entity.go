package entity

import (
	"time"
	"xyz/modules/consumer/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Limit struct {
	ID              string `json:"id"`
	LimitAmount     int    `json:"limit_amount"`
	Tenor           int    `json:"tenor"`
	ConsumerID      string `json:"consumer_id"`
	Consumer        entity.Consumer
	RemainingAmount int    `json:"remaining_amount"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (limit *Limit) BeforeCreate(*gorm.DB) (err error) {
	limit.ID = uuid.New().String()
	return
}