package request

type TopUpRequest struct {
	Balance int `json:"balance" validate:"required,max=10000000"`
}
