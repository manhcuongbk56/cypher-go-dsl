package cypher

import (
	"errors"
)

type Comparison struct {
	left     Expression
	operator Operator
	right    Expression
	key      string
	notNil   bool
	err      error
}

func ComparisonCreate(left Expression, operator Operator, right Expression) Comparison {
	if left != nil && left.getError() != nil {
		return ComparisonError(left.getError())
	}
	if operator.getError() != nil {
		return ComparisonError(operator.getError())
	}
	if right != nil && right.getError() != nil {
		return ComparisonError(right.getError())
	}
	if !operator.isUnary() {
		return ComparisonError(errors.New("operator must be unary"))
	}
	if left == nil || !left.isNotNil() {
		return ComparisonError(errors.New("left expression must be not nil"))
	}
	if right == nil || !left.isNotNil() {
		return ComparisonError(errors.New("right expression must be not nil"))
	}
	comparison := Comparison{
		left:     left,
		operator: operator,
		right:    right,
	}
	comparison.key = getAddress(&comparison)
	return comparison
}

func ComparisonCreate1(operator Operator, expression Expression) Comparison {
	if operator.getError() != nil {
		return ComparisonError(operator.getError())
	}
	if expression != nil && expression.getError() != nil {
		return ComparisonError(expression.getError())
	}
	if !operator.isUnary() {
		return ComparisonError(errors.New("comparison: operator must be unary"))
	}
	if expression == nil || !expression.isNotNil() {
		return ComparisonError(errors.New("comparison: expression must be not nil"))
	}
	var comparison Comparison
	switch operator.operatorType {
	case PREFIX:
		comparison = Comparison{
			left:     nil,
			operator: operator,
			right:    expression,
		}
	case POSTFIX:
		comparison = Comparison{
			left:     expression,
			operator: operator,
			right:    nil,
		}
	default:
		return Comparison{}
	}
	comparison.key = getAddress(&comparison)
	return comparison
}

func ComparisonError(err error) Comparison {
	return Comparison{
		err: err,
	}
}

func (c Comparison) getError() error {
	return c.err
}

func (c Comparison) isNotNil() bool {
	return c.notNil
}

func (c Comparison) accept(visitor *CypherRenderer) {
	visitor.enter(c)
	if c.left != nil {
		NameOrExpression(c.left).accept(visitor)
	}
	c.operator.accept(visitor)
	if c.right != nil {
		NameOrExpression(c.right).accept(visitor)
	}
	visitor.leave(c)
}

func (c Comparison) enter(renderer *CypherRenderer) {
}

func (c Comparison) leave(renderer *CypherRenderer) {
}

func (c Comparison) getKey() string {
	return c.key
}

func (c Comparison) GetExpressionType() ExpressionType {
	return CONDITION
}
