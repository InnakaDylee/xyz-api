package service

import (
	"errors"
	"fmt"
	"time"
	limitRepository "xyz/modules/limit/repository"
	"xyz/modules/transaction/domain"
	"xyz/modules/transaction/repository"
	"xyz/packages/validator"
)

type TransactionCommandService struct {
	transactionCommandRepository repository.TransactionCommandRepositoryInterface
	limitQueryRepository limitRepository.LimitQueryRepositoryInterface
	limitCommandRepository limitRepository.LimitCommandRepositoryInterface
}

func NewTransactionCommandService(transactionCommandRepository repository.TransactionCommandRepositoryInterface, limitQueryRepository limitRepository.LimitQueryRepositoryInterface, limitCommandRepository limitRepository.LimitCommandRepositoryInterface) TransactionCommandServiceInterface {
	return &TransactionCommandService{
		transactionCommandRepository: transactionCommandRepository,
		limitQueryRepository: limitQueryRepository,
		limitCommandRepository: limitCommandRepository,
	}
}

func (s *TransactionCommandService) CreateTransaction(transactionInput domain.Transaction) (domain.Transaction, error) {
	fmt.Printf("transaction input, ConsumerID: %v, ProductID: %v, LimitID: %v, OTR: %v, AssetName: %v ",  transactionInput.ConsumerID, transactionInput.ProductID, transactionInput.LimitID, transactionInput.OTR, transactionInput.AssetName)
	validateEmpty := validator.CheckEmpty( transactionInput.ConsumerID, transactionInput.ProductID, transactionInput.LimitID, transactionInput.OTR, transactionInput.AssetName)
	if validateEmpty != nil {
		return domain.Transaction{}, validateEmpty
	}
	fmt.Println("masuk create transaction service 1")
	limit, _ := s.limitQueryRepository.GetLimitById(transactionInput.LimitID)
	if limit.ID == "" {
		return domain.Transaction{}, errors.New("limit not found")
	}
	if limit.RemainingAmount < int(transactionInput.OTR) { 
		return domain.Transaction{}, errors.New("remaining amount is not enough")
	}
	limit.RemainingAmount -= int(transactionInput.OTR)
	if limit.RemainingAmount < 0 {
		return domain.Transaction{}, errors.New("remaining amount is not enough")
	}


	// limit, err := s.limitCommandRepository.UpdateLimit(limit)
	// if err != nil {
	// 	return domain.Transaction{}, err
	// }

	transactionInput.Limit = limit
	transactionInput.Status = "unpaid"
	transactionInput.AdminFee = 5000
	interestAmount := 0.1 * transactionInput.OTR / 12 * float64(limit.Tenor)
	transactionInput.InterestAmount = interestAmount
	installmentAmountTotal := transactionInput.OTR + interestAmount + transactionInput.AdminFee
	transactionInput.InstallmentAmount = installmentAmountTotal / float64(limit.Tenor)
	transactionInput.TransactionDate = time.Now()
	transactionInput.AmountUsed = installmentAmountTotal
	transactionInput.ContractNumber = "TRX" + time.Now().Format("20060102150405")
	transactionInput.CreatedAt = time.Now()


	transaction, err := s.transactionCommandRepository.CreateTransaction(transactionInput)
	if err != nil {
		return domain.Transaction{}, err
	}

	return transaction, nil
}