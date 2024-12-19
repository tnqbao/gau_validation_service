package mail

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau_validation/providers"
)

func SendOTPMail(c *gin.Context) {
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
	otpCode := providers.GenCaptchaCode()
	var res = providers.ResponseMail{
		To:      *req.Mail,
		Subject: "Gau OTP",
		Title:   "Verification Code",
		Body: fmt.Sprintf(`
				<div class="content">
				<p>Hello, panda deliver send you a message</p>
				<p>Use the following OTP code to complete your verification process:</p>
					<div class="otp">%s</div>
				<div>If you did not request this code, please ignore this email.</p`, otpCode),
	}
	MailSender(c, res)
	c.Set("otp", otpCode)

	c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf(`sent otp to %s`, *req.Mail), "otp": otpCode})
}
