package web

type CategoryCreateRequest struct {
	Name     string `json:"name" validate:"required,max=100,min=1"`
	IsActive bool   `json:"is_active" validate:"boolean"` // omitting validate:"required" see: https://github.com/go-playground/validator/issues/319
}
