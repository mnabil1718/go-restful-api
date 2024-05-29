package helper

import (
	"github.com/mnabil1718/go-restful-api/model/domain"
	"github.com/mnabil1718/go-restful-api/model/web"
)

// temporarily, this is as fast as we can go
func ConvertCategoriesToResponses(categories []domain.Category) []web.CategoryResponse {
	var responses []web.CategoryResponse = make([]web.CategoryResponse, len(categories))

	for idx, category := range categories {
		responses[idx] = web.CategoryResponse{
			Id:        category.Id,
			Name:      category.Name,
			IsActive:  category.IsActive,
			CreatedAt: category.CreatedAt.Time,
			UpdatedAt: category.UpdatedAt.Time,
		}
	}

	return responses
}

func ConvertCategoryToResponse(category domain.Category) web.CategoryResponse {

	return web.CategoryResponse{
		Id:        category.Id,
		Name:      category.Name,
		IsActive:  category.IsActive,
		CreatedAt: category.CreatedAt.Time,
		UpdatedAt: category.UpdatedAt.Time,
	}
}
