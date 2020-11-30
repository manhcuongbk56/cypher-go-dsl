package cypher_go_dsl

type Expression interface {
	Visitable
	GetExpressionType() ExpressionType
}

type ExpressionType string

const (
	EXPRESSION ExpressionType = "expression"
	CONDITION = "condition"
	EMPTY_CONDITION_EXPRESSION = "emptyCondition"
)
