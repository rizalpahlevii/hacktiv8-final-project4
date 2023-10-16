package service

import (
	dto "hacktiv8-final-project4/dto/product"
	"hacktiv8-final-project4/request"
)

type ProductServiceInterface interface {
	Create(request request.ProductRequest) dto.ProductDTO
	Update(id uint, request request.ProductRequest) dto.ProductDTO
	All() []dto.ProductDTO
	Delete(id uint) interface{}
}
