package app

import (
	"log"
	"restapi/config"
	"restapi/model/entity"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {

	loadConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dsn := "host=" + loadConfig.DB_HOST + " user=" + loadConfig.DB_USERNAME + " password=" + loadConfig.DB_PASSWORD + " dbname=" + loadConfig.DB_DATABASE + " port=" + strconv.Itoa(loadConfig.DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(db)
	}

	err = db.AutoMigrate(entity.Order{}, entity.Item{})

	if err != nil {
		panic(err)
	}

	return db
}
