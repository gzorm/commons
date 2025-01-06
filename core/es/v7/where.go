package v7

// ComparisonOperator 枚举定义不同的比较操作符
type ComparisonOperator string

const (
	// GreaterThan 大于
	GreaterThan ComparisonOperator = "gt"
	// GreaterThanOrEqual 大于等于
	GreaterThanOrEqual ComparisonOperator = "gte"
	// LessThan 小于
	LessThan ComparisonOperator = "lt"
	// LessThanOrEqual 小于等于
	LessThanOrEqual ComparisonOperator = "lte"
	// Equal 等于
	Equal ComparisonOperator = "eq"
	// Match 精确匹配
	Match ComparisonOperator = "match"
	// WildcardLeft 通配符左侧匹配
	WildcardLeft ComparisonOperator = "wildcard_left"
	// WildcardRight 通配符右侧匹配
	WildcardRight ComparisonOperator = "wildcard_right"
	// Wildcard 通配符匹配
	Wildcard ComparisonOperator = "wildcard"
	// In 包含
	In ComparisonOperator = "in"
	// NotIn 不包含
	NotIn ComparisonOperator = "not_in"
	// NotEqual 不等于
	NotEqual ComparisonOperator = "neq"
)

// QueryCondition 表示一个查询条件
type QueryCondition struct {
	Field    string
	Operator ComparisonOperator
	Value    interface{}
}

// SearchWithPagination 查询带分页
type SortField struct {
	Field string
	Order string
}
