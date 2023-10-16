package service

import (
	"errors"
	"github.com/go-playground/validator/v10"
	dto "hacktiv8-final-project4/dto/product"
	"hacktiv8-final-project4/exception"
	"hacktiv8-final-project4/helper"
	"hacktiv8-final-project4/repository"
	"hacktiv8-final-project4/request"
)

type ProductService struct {
	productRepository *repository.ProductRepository
	Validate          *validator.Validate
}

func NewProductService(productRepository *repository.ProductRepository, validate *validator.Validate) *ProductService {
	return &ProductService{productRepository: productRepository, Validate: validate}
}

func (service ProductService) Create(request request.ProductRequest) dto.ProductDTO {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	product := service.productRepository.Create(request)
	return dto.NewProductDTO(product)
}

func (service ProductService) Update(id uint, request request.ProductRequest) dto.ProductDTO {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	product := service.productRepository.FindById(id)
	if product.ID == 0 {
		panic(exception.NewNotFoundError(errors.New("product not found")))
	}

	product = service.productRepository.Update(product, request)
	return dto.NewProductDTO(product)
}

func (service ProductService) All() []dto.ProductDTO {
	products := service.productRepository.All()
	var productsDTO []dto.ProductDTO
	for _, value := range products {
		productsDTO = append(productsDTO, dto.NewProductDTO(value))
	}
	return productsDTO
}

func (service ProductService) Delete(id uint) error {
	product := service.productRepository.FindById(id)
	if product.ID == 0 {
		panic(exception.NewNotFoundError(errors.New("product not found")))
	}

	service.productRepository.Delete(product)
	return nil
}
