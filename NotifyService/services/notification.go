package services

import (
	"NotifyService/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func login() (string, error) {
	fmt.Println("login")
	email := os.Getenv("AUTH_SERVICE_ADMIN_LOGIN")
	password := os.Getenv("AUTH_SERVICE_ADMIN_PASSWORD")
	host := os.Getenv("AUTH_SERVICE_HOST")
	url := host + "/login"

	loginReq := LoginRequest{
		Email:    email,
		Password: password,
	}

	reqBody, err := json.Marshal(loginReq)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("auth error: %d", resp.StatusCode)
	}
	cookie, err := resp.Cookies()[0].Value, err
	if err != nil {
		return "", fmt.Errorf("cookie hollow error: %v", err)
	}
	fmt.Println("return cookie")
	return cookie, nil
}

func getUserData(userId int) map[string]interface{} {
	host := os.Getenv("AUTH_SERVICE_HOST")
	url := host + "/getUserInfo?id=" + strconv.Itoa(userId)

	// Получаем токен с помощью авторизации
	token, err := login()
	if err != nil {
		fmt.Println("auth error:", err)
		return nil
	}

	// Создаем новый GET запрос
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("creating request error:", err)
		return nil
	}

	// Добавляем токен в cookie
	req.AddCookie(&http.Cookie{
		Name:  "Authorization",
		Value: token,
	})

	// Создаем HTTP клиент и отправляем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		return nil
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error read answer:", err)
		return nil
	}

	// Декодируем JSON в map
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error pars JSON:", err)
		return nil
	}

	return result
}

func NotifyUsers(event models.TaskEvent) {
	// Собираем список уникальных получателей
	recipients := append(event.ObserversIDs, event.PerformerID, event.CreatorID)
	log.Printf("📩 Sending notification to users: %v\n", recipients)

	// Множество для хранения уникальных email'ов
	uniqueEmails := make(map[string]struct{})

	for _, userID := range recipients {
		// Получаем данные пользователя по ID
		userData := getUserData(userID)

		// Достаем email
		email, ok := userData["email"].(string)
		if !ok || email == "" {
			log.Printf("⚠️ User %d has no valid email\n", userID)
			continue
		}

		// Добавляем email в множество (дубликаты не попадут)
		uniqueEmails[email] = struct{}{}
	}

	// Формируем сообщение
	message := event.Event + " " + event.Title + " " + event.Description

	// Отправляем уведомления на уникальные email'ы
	for email := range uniqueEmails {
		err := SendEmail(email, event.Event, message)
		if err != nil {
			log.Printf("🚨 Failed to send email to %s: %v\n", email, err)
		} else {
			log.Printf("✅ Notification sent to %s\n", email)
		}
	}

	log.Printf("📢 Total unique emails notified: %d\n", len(uniqueEmails))
}
