package repository

import (
	"hacktiv8-final-project4/model"
	"hacktiv8-final-project4/request"
)

type CategoryRepositoryInterface interface {
	Create(request request.CategoryRequest) model.Category
	All() []model.Category
	FindById(id int) model.Category
	FindByType(categoryType string) model.Category
	Update(category model.Category, request request.CategoryRequest) model.Category
	Delete(category model.Category)
	IncreaseSoldProductAmount(category model.Category, amount int)
	DecreaseSoldProductAmount(category model.Category, amount int)
}
