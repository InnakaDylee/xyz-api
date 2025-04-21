package domain

import (
	"time"
	limitDomain "xyz/modules/limit/domain"
	consumerDomain "xyz/modules/consumer/domain"
	"xyz/modules/transaction/entity"
)

type Transaction struct {
	ID      string       `json:"transaction_id"`
	ConsumerID         string       `json:"Consumer_id"`
	Consumer 		   consumerDomain.Consumer
	ContractNumber     string    `json:"contract_number"`
	OTR                float64   `json:"otr"`
	AdminFee           float64   `json:"admin_fee"`
	InstallmentAmount  float64   `json:"installment_amount"`
	InterestAmount     float64   `json:"interest_amount"`
	AssetName          string    `json:"asset_name"`
	TransactionDate    time.Time `json:"transaction_date"` 
	Status             string    `json:"status"`           
	AmountUsed         float64   `json:"amount_used"`
	LimitID            string    `json:"limit_id"`
	Limit 			   limitDomain.Limit
	ProductID          int       `json:"product_id"`
	CreatedAt          time.Time `json:"created_at"` 
	UpdatedAt          time.Time `json:"updated_at"` 
}

func TransactionDomainIntoEntity(transaction Transaction) entity.Transaction {
	return entity.Transaction{
		ID:     			transaction.ID,
		ConsumerID:         transaction.ConsumerID,
		ContractNumber:     transaction.ContractNumber,
		OTR:                transaction.OTR,
		AdminFee:           transaction.AdminFee,
		InstallmentAmount:  transaction.InstallmentAmount,
		InterestAmount:     transaction.InterestAmount,
		AssetName:          transaction.AssetName,
		TransactionDate:    transaction.TransactionDate, 
		Status:             transaction.Status,           
		AmountUsed:         transaction.AmountUsed,
		LimitID:            transaction.LimitID,
		ProductID:          transaction.ProductID,
	}
}

func TransactionEntityIntoDomain(transaction entity.Transaction) Transaction {
	return Transaction{
		ID:      			transaction.ID,
		ConsumerID:         transaction.ConsumerID,
		Consumer:           consumerDomain.ConsumerEntityIntoDomain(transaction.Consumer),
		ContractNumber:     transaction.ContractNumber,
		OTR:                transaction.OTR,
		AdminFee:           transaction.AdminFee,
		InstallmentAmount:  transaction.InstallmentAmount,
		InterestAmount:     transaction.InterestAmount,
		AssetName:          transaction.AssetName,
		TransactionDate:    transaction.TransactionDate, 
		Status:             transaction.Status,           
		AmountUsed:         transaction.AmountUsed,
		LimitID:            transaction.LimitID,
		Limit:              limitDomain.LimitEntityIntoDomain(transaction.Limit),
		ProductID:          transaction.ProductID,
	}
}
