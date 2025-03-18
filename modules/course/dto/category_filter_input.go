package dto

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryFilterInput struct {
	Page      uint   `form:"page" default:"1"`
	Take      uint   `form:"take" default:"10"`
	SortBy    string `form:"sortBy" default:"id"`
	SortOrder string `form:"sortOrder" default:"asc"`
}

func (filter *CategoryFilterInput) ApplyFilter(query *gorm.DB) *gorm.DB {
	return query
}

func (filter *CategoryFilterInput) ApplyPagination(query *gorm.DB) *gorm.DB {
	desc := filter.SortOrder == "desc"
	query = query.Order(clause.OrderByColumn{Column: clause.Column{Name: filter.SortBy}, Desc: desc})
	offset := (filter.Page - 1) * filter.Take
	query = query.Offset(int(offset)).Limit(int(filter.Take))

	return query
}
