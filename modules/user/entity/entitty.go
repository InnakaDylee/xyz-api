package entity

import "time"

type User struct {
	ID       int    `gorm:"primaryKey;"`
	Email    string `gorm:"not null;unique"`
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	IsActive bool   `gorm:"default:false"`
	CreatedAt time.Time 
	UpdatedAt time.Time
}