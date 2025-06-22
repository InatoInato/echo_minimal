package config

import (
	"echo_start/internal/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		Cfg.DBHost, Cfg.DBUser, Cfg.DBPass, Cfg.DBName, Cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannto to connect to db:", err)
	}

	if err := db.AutoMigrate(&model.Car{}); err != nil {
		log.Fatal("Migration failed: ", err)
	}

	DB = db
	fmt.Println("DB Connected")
}
