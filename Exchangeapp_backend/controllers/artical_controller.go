package controllers

import (
	"errors"
	"net/http"

	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/global"
	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateArtical(ctx *gin.Context) {
	var artical model.Artical
	if err := ctx.ShouldBindJSON(&artical); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := global.DB.AutoMigrate(&artical); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := global.DB.Create(&artical).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, artical)
}

func GetArtical(ctx *gin.Context) {
	var articals []model.Artical
	if err := global.DB.Find(&articals).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, articals)
}

func GetArticalByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var artical model.Artical
	if err := global.DB.Where("id=?", id).First(&artical).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		}
		return
	}

}
