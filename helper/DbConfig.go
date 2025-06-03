package helper

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DbHost = "localhost"
	DbPort = "3308"
	DbUser = "agung"
	DbPass = "12345"
	DbName = "contact_restful_api"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(DbUser+":"+DbPass+"@tcp("+DbHost+":"+DbPort+")/"+DbName+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	//
	//db.AutoMigrate(&model.User{})
	return db, nil
}
