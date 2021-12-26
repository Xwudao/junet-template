package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Xwudao/junet"
	"github.com/Xwudao/junet/confx"
	"github.com/Xwudao/junet/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var globalDB *gorm.DB

func GetDB() *gorm.DB {
	if globalDB == nil {
		logx.Panicf("db not init")
	}
	return globalDB
}
func InitDB() *gorm.DB {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",

		confx.GetString("mysql.username"),
		confx.GetString("mysql.password"),
		confx.GetString("mysql.host"),
		confx.GetInt("mysql.port"),
		confx.GetString("mysql.dbName"),
		confx.GetString("mysql.charset"),
	)
	fmt.Println(dsn)

	gLog := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	config := &gorm.Config{}
	if confx.GetString("app.mode") == junet.Debug {
		config.Logger = gLog
	}
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		logx.Panic(err)
	}
	globalDB = db
	return db
}
