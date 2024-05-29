package domain

import (
	"database/sql"
)

type Category struct {
	Id        int64
	Name      string
	IsActive  bool
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
