package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/mnabil1718/go-restful-api/helper"
	"github.com/mnabil1718/go-restful-api/model/domain"
	"github.com/mnabil1718/go-restful-api/model/web"
	"github.com/mnabil1718/go-restful-api/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(repository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: repository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	categories := service.CategoryRepository.FindAll(ctx, tx)
	return helper.ConvertCategoriesToResponses(categories)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int64) web.CategoryResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	category := service.CategoryRepository.FindById(ctx, tx, categoryId)
	return helper.ConvertCategoryToResponse(category)
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	category := domain.Category{
		Name:     request.Name,
		IsActive: request.IsActive,
	}
	category = service.CategoryRepository.Save(ctx, tx, category)
	return helper.ConvertCategoryToResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	category := service.CategoryRepository.FindById(ctx, tx, request.Id)
	category.Name = request.Name
	category.IsActive = request.IsActive
	category.UpdatedAt.Time = time.Now()
	category = service.CategoryRepository.Update(ctx, tx, category)
	return helper.ConvertCategoryToResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int64) {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	category := service.CategoryRepository.FindById(ctx, tx, categoryId)
	service.CategoryRepository.Delete(ctx, tx, category)
}
