package seed

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func MigrateAndSeedProduct(db *gorm.DB) {
	
	if err := db.AutoMigrate(&Product{}); err != nil {
		log.Fatal("Failed to migrate Product:", err)
	}
	products := []Product{
		{
			Name:        "Kamera Digital",
			Description: "Kamera dengan resolusi tinggi cocok untuk pemula dan profesional.",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Laptop Gaming",
			Description: "Laptop dengan performa tinggi untuk kebutuhan gaming dan rendering.",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Smartphone 5G",
			Description: "Smartphone terbaru dengan jaringan 5G dan layar AMOLED.",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}


	var count int64
	db.Model(&Product{}).Count(&count)
	if count == 0 {
		if err := db.Create(&products).Error; err != nil {
			log.Fatal("Failed to seed Products:", err)
		}
	}
}
