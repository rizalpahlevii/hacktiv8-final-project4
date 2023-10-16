package httpresponse

import (
	"hacktiv8-final-project4/model"
	"time"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func NewLoginResponse(token string) LoginResponse {
	return LoginResponse{Token: token}
}

type RegisterResponse struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewRegisterResponse(user model.User) RegisterResponse {
	return RegisterResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  user.Password,
		Balance:   user.Balance,
		CreatedAt: user.CreatedAt,
	}
}
