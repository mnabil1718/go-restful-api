package web

import "time"

// needed to filter out hidden/confidential fields from DB
// in this case, nothing is filtered out
type CategoryResponse struct {
	Id        int64
	Name      string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
