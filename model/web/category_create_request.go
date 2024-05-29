package web

type CategoryCreateRequest struct {
	Name     string `json:"name" validate:"required,max=100,min=1"`
	IsActive bool   `json:"is_active" validate:"required,boolean"`
}
