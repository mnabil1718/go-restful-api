package web

type CategoryUpdateRequest struct {
	Id       int64
	Name     string
	IsActive bool
}
