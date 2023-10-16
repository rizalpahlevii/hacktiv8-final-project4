package request

type ProductRequest struct {
	Title      string `json:"title" validate:"required"`
	Price      int    `json:"price" validate:"required"`
	Stock      int    `json:"stock" validate:"required"`
	CategoryId int    `json:"category_id" validate:"required"`
}
