package dto

import "xyz/modules/transaction/domain"

type TransactionRequest struct {
	LimitID    string  `json:"limit_id" form:"limit_id" validate:"required"`
	ProductID  int     `json:"product_id" form:"product_id" validate:"required"`
	AssetName  string  `json:"asset_name" form:"asset_name" validate:"required"`
	OTR        float64 `json:"otr" form:"otr" validate:"required"`
}

type TransactionResponse struct {
	ID      		   string     `json:"id"`
	ConsumerID         string  `json:"consumer_id"`
	ContractNumber     string  `json:"contract_number"`
	OTR                float64 `json:"otr"`
	AdminFee           float64 `json:"admin_fee"`
	InstallmentAmount  float64 `json:"installment_amount"`
	InterestAmount     float64 `json:"interest_amount"`
	AssetName          string  `json:"asset_name"`
	TransactionDate    string  `json:"transaction_date"`
	Status             string  `json:"status"`
	AmountUsed         float64 `json:"amount_used"`
	LimitID            string  `json:"limit_id"`
	LimitTenor         int     `json:"limit_tenor"`
	ProductID          int     `json:"product_id"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          string  `json:"updated_at"`
}

func TransactionRequestIntoDomain(request TransactionRequest) domain.Transaction {
	return domain.Transaction{
		LimitID:   request.LimitID,
		ProductID: request.ProductID,
		AssetName: request.AssetName,
		OTR:       request.OTR,
	}
}

func TransactionDomainIntoResponse(transaction domain.Transaction) TransactionResponse {
	return TransactionResponse{
		ID:      transaction.ID,
		ConsumerID:         transaction.ConsumerID,
		ContractNumber:     transaction.ContractNumber,
		OTR:                transaction.OTR,
		AdminFee:           transaction.AdminFee,
		InstallmentAmount:  transaction.InstallmentAmount,
		InterestAmount:     transaction.InterestAmount,
		AssetName:          transaction.AssetName,
		TransactionDate:    transaction.TransactionDate.Format("2006-01-02 15:04:05"),
		Status:             transaction.Status,
		AmountUsed:         transaction.AmountUsed,
		LimitID:            transaction.LimitID,
		LimitTenor:         transaction.Limit.Tenor,
		ProductID:          transaction.ProductID,
		CreatedAt:          transaction.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:          transaction.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func TransactionDomainIntoResponseList(transactions []domain.Transaction) []TransactionResponse {
	var transactionResponses []TransactionResponse
	for _, transaction := range transactions {
		transactionResponse := TransactionDomainIntoResponse(transaction)
		transactionResponses = append(transactionResponses, transactionResponse)
	}
	return transactionResponses
}
