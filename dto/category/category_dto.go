package dto

import (
	"hacktiv8-final-project4/model"
	"time"
)

type CategoryDTO struct {
	ID                uint      `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int       `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
}

func NewCategoryDTO(category model.Category) CategoryDTO {
	return CategoryDTO{
		ID:                category.ID,
		Type:              category.Type,
		SoldProductAmount: category.SoldProductAmount,
		CreatedAt:         category.CreatedAt,
	}
}
