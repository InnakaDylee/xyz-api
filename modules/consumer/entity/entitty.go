package entity

import (
	"time"
	"xyz/modules/user/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Consumer struct {
	ID             string      `json:"id" gorm:"primaryKey"`
	User_ID        int         `json:"user_id" gorm:"not null"`
	User           entity.User `json:"user" gorm:"foreignKey:User_ID;references:ID"`
	NIK            string      `json:"nik" gorm:"not null;unique"`
	Full_Name      string      `json:"full_name" gorm:"not null"`
	Legal_Name     string      `json:"legal_name" gorm:"not null"`
	Place_Of_Birth string      `json:"place_of_birth" gorm:"not null"`
	Date_Of_Birth  string      `json:"date_of_birth" gorm:"not null"`
	Salary         int         `json:"salary" gorm:"not null"`
	Photo_KTP      string      `json:"photo_ktp" gorm:"not null"`
	Photo_Selfie   string      `json:"photo_selfie" gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (c *Consumer) BeforeCreate(*gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
