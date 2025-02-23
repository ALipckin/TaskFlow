package initializers

import (
	"TaskStorageService/models"
	"gorm.io/gorm"
	"log"
)

func SyncDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&models.Task{})
	if err != nil {
		return
	}
	log.Println("✅ Успешная миграция")
}
