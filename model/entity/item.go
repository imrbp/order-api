package entity

type Item struct {
	ItemId      int `gorm:"primaryKey"`
	ItemCode    int
	Description string
	Quantity    int
	OrderId     int
}
