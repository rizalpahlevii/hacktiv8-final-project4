package dto

import "strconv"

type TopUpDTO struct {
	Message string `json:"message"`
}

func NewTopUpDTO(balance int) TopUpDTO {
	return TopUpDTO{Message: "Your balance has been successfully updated to Rp " + strconv.Itoa(balance) + ",00"}
}
