package models

import (
	"math"
)

func (conn *Conn) Paged(pagination *Pagination, dest interface{}) *Conn {
	var totalRows int64
	conn.desc.Count(&totalRows)
	pagination.TotalRows = int(totalRows)

	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.TotalPages = totalPages

	tx := conn.desc.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Find(dest)
	return cast(tx, conn.audit)
}

// https://dev.to/rafaelgfirmino/pagination-using-gorm-scopes-3k5f
// https://gorm.io/docs/scopes.html#pagination
type Pagination struct {
    DefaultLimit int
	Limit        int
	NextPage     int
	PreviousPage int
	Page         int
	TotalRows    int
	TotalPages   int
	Searched     string
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = p.DefaultLimit
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}
