package httpresponse

import (
	"hacktiv8-final-project4/model"
	"time"
)

type CategoryResponse struct {
	ID                uint      `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int       `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
}

type CategoryListResponse struct {
	ID                uint                           `json:"id"`
	Type              string                         `json:"type"`
	SoldProductAmount int                            `json:"sold_product_amount"`
	CreatedAt         time.Time                      `json:"created_at"`
	UpdatedAt         time.Time                      `json:"updated_at"`
	Products          []CategoryListProductsResponse `json:"products"`
}

type CategoryListProductsResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Price     int       `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCategoryListResponse(category model.Category) CategoryListResponse {
	var CategoryListProductsResponses []CategoryListProductsResponse
	for _, product := range category.Products {
		CategoryListProductsResponses = append(CategoryListProductsResponses, NewCategoryListProductsResponse(product))
	}
	if len(CategoryListProductsResponses) == 0 {
		CategoryListProductsResponses = []CategoryListProductsResponse{}
	}

	return CategoryListResponse{
		ID:                category.ID,
		Type:              category.Type,
		SoldProductAmount: category.SoldProductAmount,
		CreatedAt:         category.CreatedAt,
		UpdatedAt:         category.UpdatedAt,
		Products:          CategoryListProductsResponses,
	}
}

func NewCategoryListResponses(categories []model.Category) []CategoryListResponse {
	var categoryListResponses []CategoryListResponse
	for _, category := range categories {
		categoryListResponses = append(categoryListResponses, NewCategoryListResponse(category))
	}
	return categoryListResponses
}

func NewCategoryListProductsResponse(product model.Product) CategoryListProductsResponse {
	return CategoryListProductsResponse{ID: product.ID,
		Title:     product.Title,
		Price:     product.Price,
		Stock:     product.Stock,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func NewCategoryResponse(category model.Category) CategoryResponse {
	return CategoryResponse{
		ID:                category.ID,
		Type:              category.Type,
		SoldProductAmount: category.SoldProductAmount,
		CreatedAt:         category.CreatedAt,
	}
}
