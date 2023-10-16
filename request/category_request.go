package request

type CategoryRequest struct {
	Type string `json:"type" validate:"required"`
}
