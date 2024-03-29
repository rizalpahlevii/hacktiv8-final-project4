package service

import (
	"hacktiv8-final-project4/httpresponse"
	"hacktiv8-final-project4/request"
)

type CategoryServiceContract interface {
	Create(request request.CategoryRequest) httpresponse.CategoryResponse
	All() []httpresponse.CategoryListResponse
	FindById(id int) httpresponse.CategoryResponse
	FindByType(categoryType string) httpresponse.CategoryResponse
	Update(id int, request request.CategoryRequest) httpresponse.CategoryResponse
	Delete(id int) error
}
