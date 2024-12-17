package routes

import (
	"github.com/gin-gonic/gin"
	healthcheck "github.com/tnqbao/gau_validation/api/healthcheck"
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
		}
	}
	return r
}
