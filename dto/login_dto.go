package dto

type LoginDTO struct {
	Token string `json:"token"`
}

func NewLoginDTO(token string) LoginDTO {
	return LoginDTO{Token: token}
}
