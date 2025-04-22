package entity

import (
	"time"
	"xyz/modules/transaction/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Installment struct {
	ID           string    `json:"id" gorm:"column:id"`
	TransactionID    string       `json:"transaction_id" gorm:"column:transaction_id;size:191"`
	Transaction      entity.Transaction `gorm:"foreignKey:TransactionID;references:ID"` 
	InstallmentOrder int       `json:"installment_order" gorm:"column:installment_order"`
	PaymentAmount    int       `json:"payment_amount" gorm:"column:payment_amount"`
	DueDate          time.Time `json:"due_date" gorm:"column:due_date"`
	PaymentDate      *time.Time `json:"payment_date" gorm:"column:payment_date"`
	Status           string    `json:"status" gorm:"column:status;type:enum('unpaid','paid')"`
	CreatedAt        time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (i *Installment) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.New().String()
	return
}