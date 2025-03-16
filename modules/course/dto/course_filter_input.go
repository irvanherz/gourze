package dto

import (
	"github.com/irvanherz/gourze/utils/number_filter"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CourseFilterInput struct {
	Page      uint   `form:"page" default:"1"`
	Take      uint   `form:"take" default:"10"`
	SortBy    string `form:"sortBy" default:"id"`
	SortOrder string `form:"sortOrder" default:"asc"`
	UserId    *UserIdFilter
}

type UserIdFilter struct {
	Op  number_filter.NumberFilterOperator `form:"userId.op" default:"equals"`
	Val []uint                             `form:"userId.val"`
}

func (filter *CourseFilterInput) Apply(query *gorm.DB) *gorm.DB {
	if filter.UserId != nil && filter.UserId.Val != nil {
		switch filter.UserId.Op {
		case number_filter.Equals:
			query = query.Where("user_id = ?", filter.UserId.Val)
		case number_filter.NotEquals:
			query = query.Where("user_id != ?", filter.UserId.Val)
		case number_filter.In:
			query = query.Where("user_id IN ?", filter.UserId.Val)
		case number_filter.NotIn:
			query = query.Where("user_id NOT IN ?", filter.UserId.Val)
		case number_filter.GreaterThan:
			query = query.Where("user_id > ?", filter.UserId.Val)
		case number_filter.LessThan:
			query = query.Where("user_id < ?", filter.UserId.Val)
		case number_filter.GreaterThanOrEqual:
			query = query.Where("user_id >= ?", filter.UserId.Val)
		case number_filter.LessThanOrEqual:
			query = query.Where("user_id <= ?", filter.UserId.Val)
		}
	}

	desc := filter.SortOrder == "desc"
	query = query.Order(clause.OrderByColumn{Column: clause.Column{Name: filter.SortBy}, Desc: desc})
	offset := (filter.Page - 1) * filter.Take
	query = query.Offset(int(offset)).Limit(int(filter.Take))

	return query
}
