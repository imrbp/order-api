package app

import (
	"restapi/model/entity"
	"time"
)

func Seed() error {
	db := DBInit()
	db.AutoMigrate(entity.Order{}, entity.Item{})
	orders := entity.Order{
		CustomerName: "Test1",
		OrderAt:      time.Now(),
		Items: []entity.Item{
			{
				ItemCode:    1,
				Quantity:    10,
				Description: "lalaland",
			},
			{
				ItemCode:    2,
				Quantity:    30,
				Description: "Comehere",
			},
		},
	}
	return db.Create(&orders).Error
}
