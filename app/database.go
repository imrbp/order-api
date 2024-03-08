package app

import (
	"log"
	"restapi/config"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dsn := "host=" + config.DB_HOST + " user=" + config.DB_USERNAME + " password=" + config.DB_PASSWORD + " dbname=" + config.DB_DATABASE + " port=" + strconv.Itoa(config.DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(db)
	}

	return db
}
