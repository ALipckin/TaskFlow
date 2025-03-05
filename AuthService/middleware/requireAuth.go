package middleware

import (
	"AuthService/initializers"
	"AuthService/models"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем токен из куков
		cookie, err := r.Cookie("Authorization")
		if err != nil {
			fmt.Println("Error:", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenString := cookie.Value
		fmt.Println("TokenString:", tokenString)

		// Разбираем и проверяем токен
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil {
			fmt.Println("Invalid token:", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			fmt.Println("Invalid token claims")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Проверяем срок действия токена
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			fmt.Println("Expired token")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Ищем пользователя в базе данных
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			fmt.Println("USER NOT FOUND")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Добавляем пользователя в контекст запроса
		r = r.WithContext(models.WithUser(r.Context(), &user))

		// Передаем управление следующему обработчику
		next.ServeHTTP(w, r)
	})
}
