package number_filter

type NumberFilterOperator string

const (
	Equals             NumberFilterOperator = "equals"
	NotEquals          NumberFilterOperator = "not_equals"
	In                 NumberFilterOperator = "in"
	NotIn              NumberFilterOperator = "not_in"
	GreaterThan        NumberFilterOperator = "greater_than"
	LessThan           NumberFilterOperator = "less_than"
	GreaterThanOrEqual NumberFilterOperator = "greater_than_or_equal"
	LessThanOrEqual    NumberFilterOperator = "less_than_or_equal"
)
