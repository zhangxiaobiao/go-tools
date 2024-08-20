package db

import (
	"bb/cmd/db/entity"
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var once sync.Once

func Db() *gorm.DB {
	if db != nil {
		return db
	}
	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open("bb.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatal("数据库连接失败", err.Error())
		}

		err = db.AutoMigrate(&entity.Check{})
		if err != nil {
			log.Fatal("自动迁移时发生错误", err.Error())
		}
	})
	return db
}
