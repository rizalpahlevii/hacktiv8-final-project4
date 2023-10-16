package repository

import (
	"hacktiv8-final-project4/model"
	"hacktiv8-final-project4/request"
)

type ProductRepositoryInterface interface {
	Create(request request.ProductRequest) model.Product
	Update(product model.Product, request request.ProductRequest) model.Product
	Delete(product model.Product)
	FindById(id int) model.Product
	All() []model.Product
	DecreaseStock(product model.Product, quantity int)
	IncreaseStock(product model.Product, quantity int)
}
