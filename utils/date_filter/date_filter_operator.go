package date_filter

type DateFilterOperator string

const (
	Equals         DateFilterOperator = "equals"
	NotEquals      DateFilterOperator = "not_equals"
	In             DateFilterOperator = "in"
	NotIn          DateFilterOperator = "not_in"
	GreaterThan    DateFilterOperator = "greater_than"
	GreaterOrEqual DateFilterOperator = "greater_or_equal"
	LessThan       DateFilterOperator = "less_than"
	LessOrEqual    DateFilterOperator = "less_or_equal"
)
