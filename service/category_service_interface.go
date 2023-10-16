package service

import (
	dto "hacktiv8-final-project4/dto/category"
	"hacktiv8-final-project4/request"
)

type CategoryServiceInterface interface {
	Create(request request.CategoryRequest) dto.CategoryDTO
	All() []dto.CategoryDTO
	FindById(id uint) dto.CategoryDTO
	FindByType(categoryType string) dto.CategoryDTO
	Update(id uint, request request.CategoryRequest) dto.CategoryDTO
	Delete(id uint) error
}
