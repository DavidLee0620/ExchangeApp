package controllers

import (
	"net/http"
	"time"

	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/global"
	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/model"
	"github.com/gin-gonic/gin"
)

func CreateExchangeRate(ctx *gin.Context) {
	var exchangeRate model.ExchangeRate
	if err := ctx.ShouldBindBodyWithJSON(&exchangeRate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	exchangeRate.Data = time.Now()
	if err := global.DB.AutoMigrate(&exchangeRate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	if err := global.DB.Create(&exchangeRate).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, exchangeRate)

}
