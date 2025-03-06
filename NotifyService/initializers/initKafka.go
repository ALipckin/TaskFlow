package initializers

import (
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
)

var (
	kafkaBroker string
	topic       string
	groupID     string
	Reader      *kafka.Reader
)

func InitKafka() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env: %v", err)
	}

	kafkaBroker = os.Getenv("KAFKA_HOST")
	topic = os.Getenv("KAFKA_TOPIC")
	groupID = os.Getenv("KAFKA_GROUP_ID")
	if kafkaBroker == "" {
		log.Fatal("Переменная KAFKA_HOST не задана в .env")
	}

	Reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaBroker},
		Topic:    topic,
		GroupID:  groupID,
		MaxBytes: 10e6,
	})

	log.Println("Kafka Consumer инициализирован")
}
