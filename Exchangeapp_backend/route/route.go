package route

import (
	"time"

	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/controllers"
	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, //是否允许发送cookie和证书
		MaxAge:           12 * time.Hour,
	}))
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

		api.POST("/articles", controllers.CreateArtical)
		api.GET("/articles", controllers.GetArtical)
		api.GET("/articles/:id", controllers.GetArticalByID)

		api.POST("articles/:id/like", controllers.LikeArtical)
		api.GET("/articles/:id/like", controllers.GetArticalLikes)

	}
	return r
}
