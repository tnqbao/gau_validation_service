package mail

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/tnqbao/gau_validation/config"
	"github.com/tnqbao/gau_validation/providers"
)

func CheckOTP(c *gin.Context) {
	var req providers.RequestMail
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("UserRequest binding error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}
	if req.Mail == nil || *req.Mail == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User email is required"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get OTP"})
		return
	}
	if req.Content != &otpCode {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid OTP"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "email verified"})
}
