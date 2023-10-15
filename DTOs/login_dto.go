package DTOs

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email,min=20,max=50"`
	Password string `json:"password" validate:"required"`
}
