package dto

import (
	"xyz/modules/consumer/domain"
	"xyz/modules/user/dto"
)

type (
	UpdateConsumerRequest struct {
		NIK string `json:"nik" form:"nik" validate:"required,min=16,max=16"`
		Full_Name string `json:"full_name" form:"full_name" validate:"required,min=3,max=50"`
		Legal_Name string `json:"legal_name" form:"legal_name" validate:"required,min=3,max=50"`
		Place_Of_Birth string `json:"place_of_birth" form:"place_of_birth" validate:"required,min=3,max=50"`
		Date_Of_Birth string `json:"date_of_birth" form:"date_of_birth" validate:"required"`
		Salary int `json:"salary" form:"salary" validate:"required"`
		Photo_KTP string `json:"photo_ktp" form:"photo_ktp" validate:"required"`
		Photo_Selfie string `json:"photo_selfie" form:"photo_selfie" validate:"required"`
	}

	GetConsumerResponse struct {
		ID string `json:"id"`
		User_ID int `json:"user_id"`
		NIK string `json:"nik"`
		Full_Name string `json:"full_name"`
		Legal_Name string `json:"legal_name"`
		Place_Of_Birth string `json:"place_of_birth"`
		Date_Of_Birth string `json:"date_of_birth"`
		Salary int `json:"salary"`
		Photo_KTP string `json:"photo_ktp"`
		Photo_Selfie string `json:"photo_selfie"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)

func RegisterRequestIntoConsumerDomain(request dto.RegisterRequest, userId int) (domain.Consumer) {
	return domain.Consumer{
		User_ID: userId,
		NIK: request.NIK,
		Full_Name: request.Full_Name,
		Legal_Name: request.Legal_Name,
		Place_Of_Birth: request.Place_Of_Birth,
		Date_Of_Birth: request.Date_Of_Birth,
		Salary: request.Salary,
		Photo_KTP: request.Photo_KTP,
		Photo_Selfie: request.Photo_Selfie,
	}
}

func UpdateConsumerRequestIntoConsumerDomain(request UpdateConsumerRequest, userId int) (domain.Consumer) {
	return domain.Consumer{
		User_ID: userId,
		NIK: request.NIK,
		Full_Name: request.Full_Name,
		Legal_Name: request.Legal_Name,
		Place_Of_Birth: request.Place_Of_Birth,
		Date_Of_Birth: request.Date_Of_Birth,
		Salary: request.Salary,
		Photo_KTP: request.Photo_KTP,
		Photo_Selfie: request.Photo_Selfie,
	}
}
func ConsumerDomainIntoUpdateResponse(consumer domain.Consumer) (domain.Consumer) {
	return domain.Consumer{
		ID: consumer.ID,
		User_ID: consumer.User_ID,
		NIK: consumer.NIK,
		Full_Name: consumer.Full_Name,
		Legal_Name: consumer.Legal_Name,
		Place_Of_Birth: consumer.Place_Of_Birth,
		Date_Of_Birth: consumer.Date_Of_Birth,
		Salary: consumer.Salary,
		Photo_KTP: consumer.Photo_KTP,
		Photo_Selfie: consumer.Photo_Selfie,
	}
}

func ConsumerDomainIntoGetResponse(consumer domain.Consumer) (GetConsumerResponse) {
	return GetConsumerResponse{
		ID: consumer.ID,
		User_ID: consumer.User_ID,
		NIK: consumer.NIK,
		Full_Name: consumer.Full_Name,
		Legal_Name: consumer.Legal_Name,
		Place_Of_Birth: consumer.Place_Of_Birth,
		Date_Of_Birth: consumer.Date_Of_Birth,
		Salary: consumer.Salary,
		Photo_KTP: consumer.Photo_KTP,
		Photo_Selfie: consumer.Photo_Selfie,
		CreatedAt: consumer.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: consumer.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}