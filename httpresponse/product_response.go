package httpresponse

import (
	"hacktiv8-final-project4/model"
	"time"
)

type UpdateProductResponse struct {
	Product struct {
		ID         int       `json:"id"`
		Title      string    `json:"title"`
		Price      int       `json:"price"`
		Stock      int       `json:"stock"`
		CategoryID int       `json:"category_id"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	} `json:"product"`
}

func NewUpdateProductResponse(product model.Product) UpdateProductResponse {
	return UpdateProductResponse{
		Product: struct {
			ID         int       `json:"id"`
			Title      string    `json:"title"`
			Price      int       `json:"price"`
			Stock      int       `json:"stock"`
			CategoryID int       `json:"category_id"`
			CreatedAt  time.Time `json:"created_at"`
			UpdatedAt  time.Time `json:"updated_at"`
		}{
			ID:         product.ID,
			Title:      product.Title,
			Price:      product.Price,
			Stock:      product.Stock,
			CategoryID: product.CategoryID,
			CreatedAt:  product.CreatedAt,
			UpdatedAt:  product.UpdatedAt,
		},
	}
}

type ProductResponse struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewProductResponse(product model.Product) ProductResponse {
	return ProductResponse{
		ID:         product.ID,
		Title:      product.Title,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryID: product.CategoryID,
		CreatedAt:  product.CreatedAt,
	}
}
