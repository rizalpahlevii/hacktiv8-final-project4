package repository

import (
	"gorm.io/gorm"
	"hacktiv8-final-project4/model"
	"hacktiv8-final-project4/request"
)

type CategoryRepository struct {
	*gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: DB}
}

func (repository CategoryRepository) Create(request request.CategoryRequest) model.Category {
	category := model.Category{
		Type: request.Type,
	}
	repository.DB.Create(&category)
	return category
}

func (repository CategoryRepository) All() []model.Category {
	var categories []model.Category
	repository.DB.Preload("Products").Find(&categories)
	return categories
}

func (repository CategoryRepository) FindById(id int) model.Category {
	var category model.Category
	repository.DB.Preload("Products").Where("id = ?", id).First(&category)
	return category
}

func (repository CategoryRepository) FindByType(categoryType string) model.Category {
	var category model.Category
	repository.DB.Preload("Products").Where("type = ?", categoryType).First(&category)
	return category
}

func (repository CategoryRepository) Update(category model.Category, request request.CategoryRequest) model.Category {
	category.Type = request.Type
	repository.DB.Save(&category)
	return category
}

func (repository CategoryRepository) Delete(category model.Category) {
	repository.DB.Delete(&category)
}

func (repository CategoryRepository) IncreaseSoldProductAmount(category model.Category, amount int) {
	category.SoldProductAmount += amount
	repository.DB.Save(&category)
}

func (repository CategoryRepository) DecreaseSoldProductAmount(category model.Category, amount int) {
	category.SoldProductAmount -= amount
	repository.DB.Save(&category)
}
