package middleware

import (
	"log"
	"net/http"

	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		log.Printf("AuthMiddleware: Authorization header: %s\n", token)
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"err": "Missing Authorization"})
			ctx.Abort() //出现错误，提前结束请求处理，不处理后续中间件
			return
		}
		username, err := utils.ParseJWT(token)
		if err != nil {
			log.Printf("AuthMiddleware: ParseJWT error: %v\n", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			ctx.Abort()
			return
		}
		log.Printf("AuthMiddleware: Parsed username: %s\n", username)
		ctx.Set("username", username)
		ctx.Next()
	}
}
