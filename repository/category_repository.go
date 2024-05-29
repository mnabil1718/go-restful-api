package repository

import (
	"context"
	"database/sql"

	"github.com/mnabil1718/go-restful-api/model/domain"
)

// not returning errors, because all errors will be panicked and
// handled by httprouter error handler
type CategoryRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
	FindById(ctx context.Context, tx *sql.Tx, categoryId int64) domain.Category
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
}
