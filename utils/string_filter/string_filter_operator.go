package string_filter

type StringFilterOperator string

const (
	Equals     StringFilterOperator = "equals"
	Contains   StringFilterOperator = "contains"
	StartsWith StringFilterOperator = "starts_with"
	EndsWith   StringFilterOperator = "ends_with"
	NotEquals  StringFilterOperator = "not_equals"
	In         StringFilterOperator = "in"
	NotIn      StringFilterOperator = "not_in"
)
