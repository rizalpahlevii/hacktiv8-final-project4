package model

type Category struct {
	GormModel
	Type              string    `gorm:"column:type" json:"type"`
	SoldProductAmount int       `gorm:"column:sold_product_amount" json:"sold_product_amount"`
	Products          []Product `json:"products"`
}
