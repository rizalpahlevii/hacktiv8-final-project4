package service

import (
	"errors"
	"github.com/go-playground/validator/v10"
	dto "hacktiv8-final-project4/dto/category"
	"hacktiv8-final-project4/exception"
	"hacktiv8-final-project4/helper"
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

func (service CategoryService) Create(request request.CategoryRequest) dto.CategoryDTO {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	category := service.categoryRepository.Create(request)
	return dto.NewCategoryDTO(category)
}

func (service CategoryService) All() []dto.CategoryDTO {

	categories := service.categoryRepository.All()
	var categoriesDTO []dto.CategoryDTO
	for _, value := range categories {
		categoriesDTO = append(categoriesDTO, dto.NewCategoryDTO(value))
	}
	return categoriesDTO
}

func (service CategoryService) FindById(id uint) dto.CategoryDTO {
	//TODO implement me
	panic("implement me")
}

func (service CategoryService) FindByType(categoryType string) dto.CategoryDTO {
	//TODO implement me
	panic("implement me")
}

func (service CategoryService) Update(id uint, request request.CategoryRequest) dto.CategoryDTO {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	category := service.categoryRepository.FindById(id)
	if category.ID == 0 {
		panic(exception.NewNotFoundError(errors.New("category not found")))
	}

	category = service.categoryRepository.Update(category, request)
	return dto.NewCategoryDTO(category)
}

func (service CategoryService) Delete(id uint) error {
	category := service.categoryRepository.FindById(id)
	if category.ID == 0 {
		panic(exception.NewNotFoundError(errors.New("category not found")))
	}

	service.categoryRepository.Delete(category)
	return nil
}
