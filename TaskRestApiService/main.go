package main

import (
	"TaskRestApiService/controllers"
	"TaskRestApiService/initializers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.InitProducer()
}

func main() {
	defer initializers.KafkaProducer.Close()

	mode := os.Getenv("GIN_MODE")
	gin.SetMode(mode)

	grpcClient := initializers.InitTaskStorageService()
	taskController := controllers.NewTaskController(grpcClient)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API is working",
		})
	})

	// Пример использования Kafka
	r.GET("/send-log", func(c *gin.Context) {
		err := initializers.SendMessage("logs_topic", "Тестовое сообщение в Kafka")
		if err != nil {
			log.Printf("Ошибка Kafka: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка Kafka"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Лог отправлен в Kafka"})
	})

	tasksGroup := r.Group("/tasks")
	{
		tasksGroup.POST("", taskController.TasksCreate)
		tasksGroup.GET("", taskController.TasksIndex)
		tasksGroup.GET("/:id", taskController.TasksShow)
		tasksGroup.PUT("/:id", taskController.TasksUpdate)
		tasksGroup.DELETE("/:id", taskController.TasksDelete)
	}

	authGroup := r.Group("/auth")
	{
		authHost := os.Getenv("AUTH_SERVICE_HOST")

		authGroup.POST("/login", func(c *gin.Context) {
			log.Printf("received request")
			targetURL := c.DefaultQuery("url", authHost+"/login")

			controllers.ProxyRequest(c, targetURL)
		})
		authGroup.POST("/register", func(c *gin.Context) {
			log.Printf("received request")
			targetURL := c.DefaultQuery("url", authHost+"/register")

			controllers.ProxyRequest(c, targetURL)
		})
		authGroup.GET("/validate", func(c *gin.Context) {
			log.Printf("received request")
			targetURL := c.DefaultQuery("url", authHost+"/validate")

			controllers.ProxyRequest(c, targetURL)
		})
	}

	r.Run()
}
