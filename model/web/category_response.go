package web

import "time"

// needed to filter out hidden/confidential fields from DB
// in this case, nothing is filtered out
type CategoryResponse struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
