// Package pagination TODO
package pagination

import (
	"math"

	"gorm.io/gorm"
)

// AbstractPagination 分页公共模块
type AbstractPagination[T any] struct {
	Limit int    `json:"limit,omitempty;query:limit"`
	Page  int    `json:"page,omitempty;query:page"`
	Sort  string `json:"sort,omitempty;query:sort"`
	Rows  []T    `json:"rows"`
}

// GetOffset TODO
func (p *AbstractPagination[T]) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

// GetLimit TODO
func (p *AbstractPagination[T]) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

// GetPage TODO
func (p *AbstractPagination[T]) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

// GetSort TODO
func (p *AbstractPagination[T]) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id desc"
	}
	return p.Sort
}

// SetRows TODO
func (p *AbstractPagination[T]) SetRows(rows []T) {
	p.Rows = rows
}

// SimplePagination 简单分页
type SimplePagination[T any] struct {
	*AbstractPagination[T]
	HasNextPage bool `json:"has_next_page"`
}

// SetRows 简单分页设置数据
func (s *SimplePagination[T]) SetRows(rows []T) {
	if len(rows) > s.GetLimit() {
		s.HasNextPage = true
		s.Rows = rows[:len(rows)-1]
	} else {
		s.HasNextPage = false
		s.Rows = rows
	}
}

// NewSimplePagination TODO
func NewSimplePagination[T any](limit int, page int, sort string) *SimplePagination[T] {
	return &SimplePagination[T]{
		AbstractPagination: &AbstractPagination[T]{
			Limit: limit,
			Page:  page,
			Sort:  sort,
		},
		HasNextPage: false,
	}
}

// SimplePaginate 简单分页Scope
func SimplePaginate[T any](value interface{}, pagination *SimplePagination[T]) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		limit := pagination.GetLimit() + 1
		return db.Model(value).Offset(pagination.GetOffset()).Limit(limit)
	}
}

// LengthAwarePagination 分页-统计总数
type LengthAwarePagination[T any] struct {
	*AbstractPagination[T]
	TotalRows  int64 `json:"total_rows"`
	TotalPages int   `json:"total_pages"`
}

// NewLengthAwarePagination TODO
func NewLengthAwarePagination[T any](limit int, page int, sort string) *LengthAwarePagination[T] {
	return &LengthAwarePagination[T]{
		AbstractPagination: &AbstractPagination[T]{
			Limit: limit,
			Page:  page,
			Sort:  sort,
		},
	}
}

// LengthAwarePaginate 分页Scope
// 自动完成统计总数&分页
func LengthAwarePaginate[T any](value interface{}, pagination *LengthAwarePagination[T],
	db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
