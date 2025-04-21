package domain

import (
	"time"
	"xyz/modules/consumer/entity"
)

type Consumer struct{
	ID string
	User_ID int
	NIK string
	Full_Name string
	Legal_Name string
	Place_Of_Birth string
	Date_Of_Birth string
	Salary int
	Photo_KTP string
	Photo_Selfie string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ConsumerDomainIntoEntity(consumer Consumer) (entity.Consumer) {
	return entity.Consumer{
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
		CreatedAt: consumer.CreatedAt,
		UpdatedAt: consumer.UpdatedAt,
	}
}

func ConsumerEntityIntoDomain(consumer entity.Consumer) (Consumer) {
	return Consumer{
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
		CreatedAt: consumer.CreatedAt,
		UpdatedAt: consumer.UpdatedAt,
	}
}