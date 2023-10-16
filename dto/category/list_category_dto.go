package dto

import "time"

type ListCategoryDTO struct {
	CategoryDTO
	UpdatedAt time.Time `json:"updated_at"`
}

func NewListCategoryDTO(categoryDTO CategoryDTO, updatedAt time.Time) *ListCategoryDTO {
	return &ListCategoryDTO{CategoryDTO: categoryDTO, UpdatedAt: updatedAt}
}
