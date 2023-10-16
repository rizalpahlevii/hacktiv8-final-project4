package dto

import (
	"hacktiv8-final-project4/model"
	"time"
)

type ProductDTO struct {
	ID         uint      `json:"id"`
	Title      string    `json:"name"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewProductDTO(product model.Product) ProductDTO {
	return ProductDTO{
		ID:         product.ID,
		Title:      product.Title,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryID: product.CategoryID,
		CreatedAt:  product.CreatedAt,
	}
}
