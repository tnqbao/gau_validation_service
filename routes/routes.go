package routes

import (
	"github.com/gin-gonic/gin"
	healthcheck "github.com/tnqbao/gau_validation/api/healthcheck"
	mail "github.com/tnqbao/gau_validation/api/mail"
	"github.com/tnqbao/gau_validation/middlewares"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	apiRoutes := r.Group("/api")
	{
		validateRoutes := apiRoutes.Group("/validation")
		{
			validateRoutes.GET("/check", healthcheck.Healthcheck)

			mailRoutes := validateRoutes.Group("/mail")
			{
				mailRoutes.POST("/get-otp", mail.SendOTPMail)
				mailRoutes.POST("/check-otp", mail.CheckOTP)
			}
		}
	}
	return r
}
