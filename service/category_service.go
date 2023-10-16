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

type CategoryService struct {
	categoryRepository *repository.CategoryRepository
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository *repository.CategoryRepository, validate *validator.Validate) *CategoryService {
	return &CategoryService{categoryRepository: categoryRepository, Validate: validate}
}

func (service CategoryService) Create(request request.CategoryRequest) httpresponse.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	category := service.categoryRepository.Create(request)
	return httpresponse.NewCategoryResponse(category)
}

func (service CategoryService) All() []httpresponse.CategoryListResponse {
	categories := service.categoryRepository.All()
	return httpresponse.NewCategoryListResponses(categories)
}

func (service CategoryService) FindById(id int) httpresponse.CategoryResponse {
	//TODO implement me
	panic("implement me")
}

func (service CategoryService) FindByType(categoryType string) httpresponse.CategoryResponse {
	//TODO implement me
	panic("implement me")
}

func (service CategoryService) Update(id int, request request.CategoryRequest) httpresponse.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	category := service.categoryRepository.FindById(id)
	if category.ID == 0 {
		panic(exception.NewNotFoundError(errors.New("category not found")))
	}

	category = service.categoryRepository.Update(category, request)
	return httpresponse.NewCategoryResponse(category)
}

func (service CategoryService) Delete(id int) error {
	category := service.categoryRepository.FindById(id)
	if category.ID == 0 {
		panic(exception.NewNotFoundError(errors.New("category not found")))
	}

	service.categoryRepository.Delete(category)
	return nil
}
