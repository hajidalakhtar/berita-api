package config

import (
	"bebasinfo/exception"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(configuration Config) *gorm.DB {

	dsn := configuration.Get("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	exception.PanicIfNeeded(err)
	//db.AutoMigrate(&entity.News{})

	return db

}
