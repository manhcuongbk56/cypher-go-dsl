package cypher

type Expression interface {
	Visitable
	GetExpressionType() ExpressionType
}

type ExpressionType string

const (
	EXPRESSION                 ExpressionType = "expression"
	CONDITION                                 = "conditionBuilder"
	EMPTY_CONDITION_EXPRESSION                = "emptyCondition"
	LITERAL                                   = "literal"
)
