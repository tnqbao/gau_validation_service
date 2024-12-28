package mail

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau_validation_service/providers"
	gomail "gopkg.in/mail.v2"
)

func MailSender(c *gin.Context, res providers.ResponseMail) {
	htmlBody := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>%s</title>
			<style>
				body { font-family: Arial, sans-serif; margin: 0; padding: 0; background-color: #f4f4f4;}
				.container { max-width: 600px; margin: 0 auto; background: #fff; border-radius: 8px; box-shadow: 0 4px 8px rgba(0,0,0,0.1); border: 0.5rem solid;}
				.header { background: #007bff; color: #fff; padding: 10px; text-align: center; }
				.content { background: #ffffff; color: #000000; padding-top: 10px ; padding-left: 10px; padding-right: 10px;text-align: center; }
				.footer { background: #ffffff; color: #000000; padding: 10px; text-align: center; font-size: 14px; border-bottom-left-radius: 8px; border-bottom-right-radius: 8px; }
				.otp { font-size: 24px; font-weight: bold; margin: 20px 0; text-align: center; color: #007bff; }
				.text-center { text-align: center; }
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<h1>%s</h1>
				</div>
				%s
				<div class="footer">
					<p>Thank you for using our service!</p>
					<p>&copy; 2024 Gau. All rights reserved.</p>
				</div>
			</div>
		</body>
		</html>
	`, res.Title, res.Title, res.Body)

	message := gomail.NewMessage()

	message.SetHeader("From", os.Getenv("SMTP_USERNAME"))
	message.SetHeader("To", res.To)
	message.SetHeader("Subject", res.Subject)

	message.SetBody("text/html", htmlBody)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"))

	if err := dialer.DialAndSend(message); err != nil {
		log.Println("Error sending email:", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to send email"})
		return
	}

	c.Next()
}
