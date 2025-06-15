package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", func(ctx *gin.Context) {
			// login logic
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"message": "login success",
			})
		})
		auth.POST("/register", func(ctx *gin.Context) {
			// register logic
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"message": "register success",
			})
		})
	}
	return r
}
