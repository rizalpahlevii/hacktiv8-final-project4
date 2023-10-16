package httpresponse

import "strconv"

type TopUpResponse struct {
	Message string `json:"message"`
}

func NewTopUpResponse(balance int) TopUpResponse {
	return TopUpResponse{Message: "Your balance has been successfully updated to Rp " + strconv.Itoa(balance) + ",00"}
}
