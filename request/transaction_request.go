package request

type TransactionRequest struct {
	ProductId  int `json:"product_id" validate:"required"`
	Quantity   int `json:"quantity" validate:"required"`
	UserId     int
	TotalPrice int
}
