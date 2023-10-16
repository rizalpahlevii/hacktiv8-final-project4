package repository

import (
	"gorm.io/gorm"
	"hacktiv8-final-project4/model"
	"hacktiv8-final-project4/request"
)

type ProductRepository struct {
	*gorm.DB
}

func NewProductRepository(DB *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: DB}
}

func (repository ProductRepository) Create(request request.ProductRequest) model.Product {
	product := model.Product{
		Title:      request.Title,
		Price:      request.Price,
		Stock:      request.Stock,
		CategoryID: request.CategoryId,
	}
	repository.DB.Create(&product)
	return product
}

func (repository ProductRepository) FindById(id uint) model.Product {
	var product model.Product
	repository.DB.First(&product, id)
	return product
}

func (repository ProductRepository) Update(product model.Product, request request.ProductRequest) model.Product {
	product.Title = request.Title
	product.Price = request.Price
	product.Stock = request.Stock
	product.CategoryID = request.CategoryId
	repository.DB.Save(&product)
	return product
}

func (repository ProductRepository) Delete(product model.Product) {
	repository.DB.Delete(&product)
}

func (repository ProductRepository) All() []model.Product {
	var products []model.Product
	repository.DB.Preload("Category").Find(&products)
	return products
}
