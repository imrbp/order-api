package entity

import "time"

type Order struct {
	Id           int    `gorm:"primaryKey"`
	CustomerName string `gorm:"type:varchar(50)"`
	OrderAt      time.Time
	Items        []Item
}
