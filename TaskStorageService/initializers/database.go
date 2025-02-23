package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDB() *gorm.DB {
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("❌ Ошибка: DATABASE_URL не установлен")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database")
	}
	log.Println("✅ Успешное подключение к базе данных")
	return db
}
