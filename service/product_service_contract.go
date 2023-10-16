package service

import (
	"hacktiv8-final-project4/httpresponse"
	"hacktiv8-final-project4/request"
)

type ProductServiceContract interface {
	Create(request request.ProductRequest) httpresponse.ProductResponse
	Update(id int, request request.ProductRequest) httpresponse.UpdateProductResponse
	All() []httpresponse.ProductResponse
	Delete(id int) error
}
