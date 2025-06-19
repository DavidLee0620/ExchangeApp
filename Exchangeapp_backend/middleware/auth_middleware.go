package middleware

import (
	"net/http"

	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"err": "Missing Authorization"})
			ctx.Abort() //出现错误，提前结束请求处理，不处理后续中间件
			return
		}
		uesrname, err := utils.ParseJWT(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invaild Token"})
			ctx.Abort()
			return
		}
		ctx.Set("username", uesrname)
		ctx.Next()
	}
}
