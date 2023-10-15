package model

type Product struct {
	GormModel
	Title      string   `gorm:"column:title" json:"title"`
	Price      int      `gorm:"column:price" json:"price"`
	Stock      int      `gorm:"column:stock" json:"stock"`
	CategoryID int      `gorm:"column:category_id" json:"category_id"`
	Category   Category `json:"category"`
}
