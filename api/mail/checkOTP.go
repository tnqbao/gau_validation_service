package mail

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/tnqbao/gau_validation_service/config"
	"github.com/tnqbao/gau_validation_service/providers"
)

func CheckOTP(c *gin.Context) {
	var req providers.RequestMail
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Request binding error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}

	if req.Mail == nil || *req.Mail == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User email is required"})
		return
	}

	if req.Content == nil || *req.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "OTP content is required"})
		return
	}

	ctx := context.Background()
	redisClient := config.GetRedisClient()
	key := "otp:" + *req.Mail

	otpCode, err := redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "OTP not found"})
		return
	} else if err != nil {
		log.Printf("Failed to get OTP from Redis: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if *req.Content != otpCode {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OTP"})
		return
	}

	if err := providers.UpdateUserBooleanField("is_email_verified", true, c); err != nil {
		log.Printf("Failed to update user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email verified"})
}
