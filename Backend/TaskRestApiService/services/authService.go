package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func GetUsersData(ids []int) ([]User, error) {
	authServiceURL := os.Getenv("AUTH_SERVICE_URL")
	authToken := os.Getenv("AUTH_SERVICE_TOKEN")
	if len(ids) == 0 {
		return nil, fmt.Errorf("missing ids parameter")
	}

	idStrings := make([]string, len(ids))
	for i, id := range ids {
		idStrings[i] = strconv.Itoa(id)
	}

	url := fmt.Sprintf("%s/users?ids=%s", authServiceURL, strings.Join(idStrings, ","))
	log.Printf("Request URL: %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, fmt.Errorf("failed to create request")
	}

	req.Header.Set("Authorization", authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error fetching users data: %v", err)
		return nil, fmt.Errorf("failed to fetch users data")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch users data, status: %d", resp.StatusCode)
	}

	var responseData []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return nil, fmt.Errorf("failed to parse response")
	}

	return ParseUsersData(responseData)
}

func ParseUsersData(data []map[string]interface{}) ([]User, error) {
	var users []User

	for _, item := range data {
		id, ok := item["id"].(float64)
		if !ok {
			return nil, fmt.Errorf("invalid id format")
		}
		email, ok := item["email"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid email format")
		}
		name, ok := item["name"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid name format")
		}
		users = append(users, User{
			ID:    int(id),
			Email: email,
			Name:  name,
		})
	}

	return users, nil
}

func ValidateToken(token string) (User, error) {
	authServiceURL := os.Getenv("AUTH_SERVICE_URL")

	var user User

	if token == "" {
		return user, fmt.Errorf("no token provided")
	}

	if !isBearerToken(token) {
		return user, fmt.Errorf("invalid token format")
	}

	client := http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", authServiceURL+"/validate", nil)
	if err != nil {
		return user, fmt.Errorf("internal server error: %w", err)
	}

	req.Header.Set("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		return user, fmt.Errorf("auth service error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return user, fmt.Errorf("invalid token: status code %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return user, fmt.Errorf("failed to parse auth service response: %w", err)
	}

	return user, nil
}

func isBearerToken(token string) bool {
	return len(token) > 7 && token[:7] == "Bearer "
}
