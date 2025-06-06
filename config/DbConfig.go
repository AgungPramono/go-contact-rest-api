package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
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

	// Ambil instance SQL dari GORM
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Setting connection pool
	sqlDB.SetMaxOpenConns(80)                  // Maks koneksi aktif
	sqlDB.SetMaxIdleConns(25)                  // Maks koneksi idle
	sqlDB.SetConnMaxIdleTime(10 * time.Minute) // Waktu koneksi idle sebelum dibuang
	sqlDB.SetConnMaxLifetime(1 * time.Hour)    // Waktu maksimal hidupnya koneksi
	//
	//db.AutoMigrate(&model.User{}, &model.Contact{}, &model.Address{})

	return db, nil
}
