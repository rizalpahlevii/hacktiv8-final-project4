package service

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"hacktiv8-final-project4/exception"
	"hacktiv8-final-project4/helper"
	"hacktiv8-final-project4/httpresponse"
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

func (service ProductService) Create(request request.ProductRequest) httpresponse.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	product := service.productRepository.Create(request)
	return httpresponse.NewProductResponse(product)
}

func (service ProductService) Update(id int, request request.ProductRequest) httpresponse.UpdateProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	product := service.productRepository.FindById(id)
	if product.ID == 0 {
		panic(exception.NewNotFoundError(errors.New("product not found")))
	}

	product = service.productRepository.Update(product, request)
	return httpresponse.NewUpdateProductResponse(product)
}

func (service ProductService) All() []httpresponse.ProductResponse {
	products := service.productRepository.All()
	var productResponses []httpresponse.ProductResponse
	for _, value := range products {
		productResponses = append(productResponses, httpresponse.NewProductResponse(value))
	}
	return productResponses
}

func (service ProductService) Delete(id int) error {
	product := service.productRepository.FindById(id)
	if product.ID == 0 {
		panic(exception.NewNotFoundError(errors.New("product not found")))
	}

	service.productRepository.Delete(product)
	return nil
}
