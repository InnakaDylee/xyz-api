package dto

import "xyz/modules/limit/domain"

type (
	LimitListResponse struct {
		LimitId      string `json:"limit_id"`
		LimitAmount  int    `json:"limit_amount"`
		Tenor        int    `json:"tenor"`
		ConsumerId   string `json:"consumer_id"`
		RemainingAmount int `json:"remaining_amount"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
	}
)

func LimitDomainIntoListResponse(limit []domain.Limit) ([]LimitListResponse) {
	var limitListResponse []LimitListResponse
	for _, limit := range limit {
		limitListResponse = append(limitListResponse, LimitListResponse{
			LimitId:        limit.ID,
			LimitAmount:    limit.LimitAmount,
			Tenor:          limit.Tenor,
			ConsumerId:     limit.ConsumerID,
			RemainingAmount: limit.RemainingAmount,
			CreatedAt:      limit.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:      limit.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return limitListResponse
}