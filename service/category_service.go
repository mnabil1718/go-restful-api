package service

import (
	"context"

	"github.com/mnabil1718/go-restful-api/model/web"
)

type CategoryService interface {
	FindAll(ctx context.Context) []web.CategoryResponse
	FindById(ctx context.Context, categoryId int64) web.CategoryResponse
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int64)
}
