package httpresponse

import (
	"hacktiv8-final-project4/model"
	"strconv"
	"time"
)

type UserResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserResponse(user model.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FullName:  user.FullName,
		Balance:   user.Balance,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type TopUpResponse struct {
	Message string `json:"message"`
}

func NewTopUpResponse(balance int) TopUpResponse {
	return TopUpResponse{Message: "Your balance has been successfully updated to Rp " + strconv.Itoa(balance) + ",00"}
}
