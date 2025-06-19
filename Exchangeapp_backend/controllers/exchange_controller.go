package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/global"
	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func GetExchangeRate(ctx *gin.Context) {
	var exchangeRate []model.ExchangeRate

	if err := global.DB.Find(&exchangeRate).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		}

		return
	}
	ctx.JSON(http.StatusOK, exchangeRate)
}
