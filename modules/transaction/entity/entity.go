package entity

import (
	"time"
	consumerEntity "xyz/modules/consumer/entity"
	limitEntity "xyz/modules/limit/entity"
	"xyz/storages/seed"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID                 string             `json:"id" gorm:"primaryKey"`
	ConsumerID         string             `json:"consumer_id" gorm:"size:191"`
	Consumer           consumerEntity.Consumer `gorm:"foreignKey:ConsumerID;references:ID"` // relasi ke consumer

	ContractNumber     string             `json:"contract_number"`
	OTR                float64            `json:"otr"`
	AdminFee           float64            `json:"admin_fee"`
	InstallmentAmount  float64            `json:"installment_amount"`
	InterestAmount     float64            `json:"interest_amount"`
	AssetName          string             `json:"asset_name"`
	TransactionDate    time.Time          `json:"transaction_date"`
	Status             string             `json:"status" gorm:"type:enum('unpaid','paid')"`
	AmountUsed         float64            `json:"amount_used"`

	LimitID            string             `json:"limit_id" gorm:"size:191"`
	Limit              limitEntity.Limit  `gorm:"foreignKey:LimitID;references:ID"` // relasi ke limit

	ProductID          int                `json:"product_id"`
	Product            seed.Product            `gorm:"foreignKey:ProductID;references:ID"` // relasi ke product

	CreatedAt          time.Time          `json:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at"`
}


func (t *Transaction) BeforeCreate(*gorm.DB) (err error) {
	t.ID = uuid.New().String()
	return
}