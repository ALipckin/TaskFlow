package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func RequireAuth(c *gin.Context) {
	authServiceURL := os.Getenv("AUTH_SERVICE_URL")

	_, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No token"})
		return
	}

	client := http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("GET", authServiceURL+"/validate", nil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	for _, cookie := range c.Request.Cookies() {
		req.AddCookie(cookie)
	}

	resp, err := client.Do(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Auth service error"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
		return
	}

	c.Next()
}
