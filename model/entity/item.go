package entity

type Item struct {
	Id          int    `gorm:"primaryKey"`
	Code        string `gorm:"type:varchar(10)"`
	Description string `gorm:"type:varchar(50)"`
	Quantity    int
	OrderId     int
}
