package entity

import "time"

type Order struct {
	OrderId      int `gorm:"primaryKey"`
	CustomerName string
	OrderAt      time.Time
	Items        []Item `gorm:"foreignKey:OrderId"`
}
