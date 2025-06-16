package controllers

import (
	"net/http"

	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/model"
	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/utils"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var user model.Uesr
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.Password = hashedPwd

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})

}
