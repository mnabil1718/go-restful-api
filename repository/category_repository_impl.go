package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mnabil1718/go-restful-api/helper"
	"github.com/mnabil1718/go-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name, is_active, created_at, updated_at FROM categories"
	categories := []domain.Category{}
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	for rows.Next() {
		var category domain.Category
		err = rows.Scan(&category.Id, &category.Name, &category.IsActive, &category.CreatedAt, &category.UpdatedAt)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int64) (domain.Category, error) {
	SQL := "SELECT id, name, is_active, created_at, updated_at FROM categories WHERE id = $1"
	category := domain.Category{}
	err := tx.QueryRowContext(ctx, SQL, categoryId).Scan(&category.Id, &category.Name, &category.IsActive, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return category, errors.New("category not found")
	}
	return category, nil
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO categories (name, is_active) VALUES ($1, $2) RETURNING id, name, is_active, created_at, updated_at"
	err := tx.QueryRowContext(ctx, SQL, category.Name, category.IsActive).Scan(&category.Id, &category.Name, &category.IsActive, &category.CreatedAt, &category.UpdatedAt)
	helper.PanicIfError(err)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE categories SET name=$1,is_active=$2,updated_at=$3 WHERE id=$4 RETURNING id, name, is_active, created_at, updated_at"
	err := tx.QueryRowContext(ctx, SQL, category.Name, category.IsActive, category.UpdatedAt, category.Id).Scan(&category.Id, &category.Name, &category.IsActive, &category.CreatedAt, &category.UpdatedAt)
	helper.PanicIfError(err)
	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE FROM categories WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}
