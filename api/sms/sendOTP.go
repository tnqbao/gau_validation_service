package sms

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau_validation_service/config"
	"github.com/tnqbao/gau_validation_service/providers"
)

func SendOTPSMS(c *gin.Context) {
	var req providers.RequestSMS
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("UserRequest binding error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}
	if req.Phone == nil || *req.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone is required"})
		return
	}
	otpCode := providers.GenCaptchaCode()

	content := "Your OTP code is: " + otpCode

	sid, err := providers.SendSMSWithTwilio(*req.Phone, content)
	if err != nil {
		log.Println("Failed to send SMS:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP SMS"})
		return
	}

	log.Printf("SMS sent successfully, SID: %s", sid)

	ctx := context.Background()
	redisClient := config.GetRedisClient()
	err = redisClient.Set(ctx, "otp:"+*req.Phone, otpCode, 30*time.Minute).Err()
	if err != nil {
		log.Println("Failed to set OTP in Redis:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP SMS sent successfully"})
}
