package models

type TransactionHistory struct {
	GormModel
	ProductID  int     `gorm:"column:product_id" json:"product_id"`
	UserID     int     `gorm:"column:user_id" json:"user_id"`
	Quantity   int     `gorm:"column:quantity" json:"quantity"`
	TotalPrice int     `gorm:"column:total_price" json:"total_price"`
	User       User    `json:"user"`
	Product    Product `json:"product"`
}
