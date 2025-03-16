package dto

import (
	"github.com/irvanherz/gourze/utils/string_filter"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserFilterInput struct {
	Page      uint   `form:"page" default:"1"`
	Take      uint   `form:"take" default:"10"`
	SortBy    string `form:"sortBy" default:"id"`
	SortOrder string `form:"sortOrder" default:"asc"`
	Username  *UsernameFilter
	Email     *EmailFilter
	FullName  *FullNameFilter
}

type UsernameFilter struct {
	Op  string_filter.StringFilterOperator `form:"username.op" default:"equals"`
	Val []string                           `form:"username.val"`
}

type EmailFilter struct {
	Op  string_filter.StringFilterOperator `form:"email.op" default:"equals"`
	Val []string                           `form:"email.val"`
}

type FullNameFilter struct {
	Op  string_filter.StringFilterOperator `form:"fullName.op" default:"equals"`
	Val []string                           `form:"fullName.val"`
}

func (filter *UserFilterInput) ApplyFilter(query *gorm.DB) *gorm.DB {
	if filter.Username != nil && filter.Username.Val != nil {
		switch filter.Username.Op {
		case "equals":
			query = query.Where("username = ?", filter.Username.Val)
		case "contains":
			query = query.Where("username LIKE ?", "%"+filter.Username.Val[0]+"%")
		case "starts_with":
			query = query.Where("username LIKE ?", filter.Username.Val[0]+"%")
		case "ends_with":
			query = query.Where("username LIKE ?", "%"+filter.Username.Val[0])
		case "not_equals":
			query = query.Where("username != ?", filter.Username.Val)
		case "in":
			query = query.Where("username IN ?", filter.Username.Val)
		case "not_in":
			query = query.Where("username NOT IN ?", filter.Username.Val)
		}
	}

	if filter.Email != nil && filter.Email.Val != nil {
		switch filter.Email.Op {
		case "equals":
			query = query.Where("email = ?", filter.Email.Val)
		case "contains":
			query = query.Where("email LIKE ?", "%"+filter.Email.Val[0]+"%")
		case "starts_with":
			query = query.Where("email LIKE ?", filter.Email.Val[0]+"%")
		case "ends_with":
			query = query.Where("email LIKE ?", "%"+filter.Email.Val[0])
		case "not_equals":
			query = query.Where("email != ?", filter.Email.Val)
		case "in":
			query = query.Where("email IN ?", filter.Email.Val)
		case "not_in":
			query = query.Where("email NOT IN ?", filter.Email.Val)
		}
	}

	if filter.FullName != nil && filter.FullName.Val != nil {
		switch filter.FullName.Op {
		case "equals":
			query = query.Where("full_name = ?", filter.FullName.Val)
		case "contains":
			query = query.Where("full_name LIKE ?", "%"+filter.FullName.Val[0]+"%")
		case "starts_with":
			query = query.Where("full_name LIKE ?", filter.FullName.Val[0]+"%")
		case "ends_with":
			query = query.Where("full_name LIKE ?", "%"+filter.FullName.Val[0])
		case "not_equals":
			query = query.Where("full_name != ?", filter.FullName.Val)
		case "in":
			query = query.Where("full_name IN ?", filter.FullName.Val)
		case "not_in":
			query = query.Where("full_name NOT IN ?", filter.FullName.Val)
		}
	}
	return query
}

func (filter *UserFilterInput) ApplyPagination(query *gorm.DB) *gorm.DB {
	desc := filter.SortOrder == "desc"
	query = query.Order(clause.OrderByColumn{Column: clause.Column{Name: filter.SortBy}, Desc: desc})
	offset := (filter.Page - 1) * filter.Take
	query = query.Offset(int(offset)).Limit(int(filter.Take))

	return query
}
