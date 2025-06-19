package route

import (
	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/controllers"
	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	api := r.Group("/api")
	api.GET("/exchangeRates", controllers.GetExchangeRate)
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/exchangeRates", controllers.CreateExchangeRate)
	}
	return r
}
