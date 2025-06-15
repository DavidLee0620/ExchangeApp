package config

import (
	"log"
	"time"

	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() {
	dsn := AppConfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm.Open is err:%v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("db.DB is err:%v", err)
	}
	global.DB = db
	sqlDB.SetConnMaxIdleTime(time.Duration(AppConfig.Database.MaxIdleConns))
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

}
