package controllers

import (
	"net/http"

	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/global"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func LikeArtical(ctx *gin.Context) {
	articalID := ctx.Param("id")
	likeKey := "artical:" + articalID + ":likes" // Redis key
	if err := global.RedisDB.Incr(likeKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successful"})

}

func GetArticalLikes(ctx *gin.Context) {
	articalID := ctx.Param("id")
	likeKey := "artical:" + articalID + ":likes" // Redis key
	likes, err := global.RedisDB.Get(likeKey).Result()
	if err == redis.Nil {
		likes = "0"
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"likes": likes})
}
