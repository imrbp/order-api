package app

import (
	"restapi/model/entity"
	"time"
)

func Seed() error {
	db := DBInit()
	err := db.AutoMigrate(entity.Order{}, entity.Item{})
	if err != nil {
		return err
	}
	orders := entity.Order{
		CustomerName: "Test1",
		OrderAt:      time.Now(),
		Items: []entity.Item{
			{
				Code:        "Item1",
				Quantity:    10,
				Description: "lalaland",
			},
			{
				Code:        "item2",
				Quantity:    30,
				Description: "Comehere",
			},
		},
	}
	return db.Create(&orders).Error
}
