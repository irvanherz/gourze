package dto

import (
	"github.com/irvanherz/gourze/utils/number_filter"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderFilterInput struct {
	Page      uint   `form:"page" default:"1"`
	Take      uint   `form:"take" default:"10"`
	SortBy    string `form:"sortBy" default:"id"`
	SortOrder string `form:"sortOrder" default:"asc"`
	UserId    *UserIdFilter
	Amount    *AmountFilter
}

type UserIdFilter struct {
	Op  number_filter.NumberFilterOperator `form:"userId.op" default:"equals"`
	Val []uint                             `form:"userId.val"`
}

type AmountFilter struct {
	Op  number_filter.NumberFilterOperator `form:"amount.op" default:"equals"`
	Val []uint                             `form:"amount.val"`
}

func (filter *OrderFilterInput) ApplyFilter(query *gorm.DB) *gorm.DB {
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

	if filter.Amount != nil && filter.Amount.Val != nil {
		switch filter.Amount.Op {
		case number_filter.Equals:
			query = query.Where("amount = ?", filter.Amount.Val)
		case number_filter.NotEquals:
			query = query.Where("amount != ?", filter.Amount.Val)
		case number_filter.In:
			query = query.Where("amount IN ?", filter.Amount.Val)
		case number_filter.NotIn:
			query = query.Where("amount NOT IN ?", filter.Amount.Val)
		case number_filter.GreaterThan:
			query = query.Where("amount > ?", filter.Amount.Val)
		case number_filter.LessThan:
			query = query.Where("amount < ?", filter.Amount.Val)
		case number_filter.GreaterThanOrEqual:
			query = query.Where("amount >= ?", filter.Amount.Val)
		case number_filter.LessThanOrEqual:
			query = query.Where("amount <= ?", filter.Amount.Val)
		}
	}
	return query
}

func (filter *OrderFilterInput) ApplyPagination(query *gorm.DB) *gorm.DB {
	desc := filter.SortOrder == "desc"
	query = query.Order(clause.OrderByColumn{Column: clause.Column{Name: filter.SortBy}, Desc: desc})
	offset := (filter.Page - 1) * filter.Take
	query = query.Offset(int(offset)).Limit(int(filter.Take))

	return query
}
