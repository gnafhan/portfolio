package middleware

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func ApiKey(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	req_api_key := c.Request.Header.Get("api_key")

	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		log.Fatal("API_KEY is not set")
	}

	if req_api_key != apiKey {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
}
