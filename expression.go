package cypher

type Expression interface {
	Visitable
	GetExpressionType() ExpressionType
	As(alias string) ExpressionContainer
	IsEqualTo(rhs Expression) ConditionContainer
	IsNotEqualTo(rhs Expression) ConditionContainer
	Lt(rhs Expression) ConditionContainer
	Lte(rhs Expression) ConditionContainer
	Gt(rhs Expression) ConditionContainer
	Gte(rhs Expression) ConditionContainer
	IsTrue() ConditionContainer
	IsFalse() ConditionContainer
	Matches(expression Expression) ConditionContainer
	MatchesPattern(pattern string) ConditionContainer
	StartWiths(rhs Expression) ConditionContainer
	Contains(rhs Expression) ConditionContainer
	EndsWith(rhs Expression) ConditionContainer
	Concat(rhs Expression) ExpressionContainer
	Add(rhs Expression) ExpressionContainer
	Subtract(rhs Expression) ExpressionContainer
	Multiply(rhs Expression) ExpressionContainer
	Divide(rhs Expression) ExpressionContainer
	Remainder(rhs Expression) ExpressionContainer
	Pow(rhs Expression) ExpressionContainer
	IsNull() ConditionContainer
	IsNotNull() ConditionContainer
	In(haystack Expression) ConditionContainer
	IsEmpty() ConditionContainer
	Descending() SortItem
	Ascending() SortItem
}

type ExpressionType string

const (
	EXPRESSION                 ExpressionType = "expression"
	CONDITION                                 = "conditionBuilder"
	EMPTY_CONDITION_EXPRESSION                = "emptyCondition"
	LITERAL                                   = "literal"
)
