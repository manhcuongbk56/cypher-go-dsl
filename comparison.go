package cypher

import (
	"errors"
)

type Comparison struct {
	ConditionContainer
	left     Expression
	operator Operator
	right    Expression
	key      string
	notNil   bool
	err      error
}

func ComparisonCreate(left Expression, operator Operator, right Expression) Comparison {
	if left != nil && left.GetError() != nil {
		return ComparisonError(left.GetError())
	}
	if operator.GetError() != nil {
		return ComparisonError(operator.GetError())
	}
	if right != nil && right.GetError() != nil {
		return ComparisonError(right.GetError())
	}
	if !operator.isNotNil() {
		return ComparisonError(errors.New("operator must be not nil"))
	}
	if left == nil || !left.isNotNil() {
		return ComparisonError(errors.New("left expression must be not nil"))
	}
	if right == nil || !left.isNotNil() {
		return ComparisonError(errors.New("right expression must be not nil"))
	}
	comparison := Comparison{
		left:     nestedIfCondition(left),
		operator: operator,
		right:    nestedIfCondition(right),
		notNil:   true,
	}
	comparison.key = getAddress(&comparison)
	comparison.ConditionContainer = ConditionWrap(comparison)
	return comparison
}

func ComparisonCreate1(operator Operator, expression Expression) Comparison {
	if operator.GetError() != nil {
		return ComparisonError(operator.GetError())
	}
	if expression != nil && expression.GetError() != nil {
		return ComparisonError(expression.GetError())
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
			right:    nestedIfCondition(expression),
		}
	case POSTFIX:
		comparison = Comparison{
			left:     nestedIfCondition(expression),
			operator: operator,
			right:    nil,
		}
	default:
		return Comparison{}
	}
	comparison.key = getAddress(&comparison)
	comparison.notNil = true
	comparison.ConditionContainer = ConditionWrap(comparison)
	return comparison
}

func ComparisonError(err error) Comparison {
	return Comparison{
		err: err,
	}
}

func (c Comparison) GetError() error {
	return c.err
}

func (c Comparison) isNotNil() bool {
	return c.notNil
}

func (c Comparison) accept(visitor *CypherRenderer) {
	visitor.enter(c)
	if c.left != nil && c.left.isNotNil() {
		NameOrExpression(c.left).accept(visitor)
	}
	c.operator.accept(visitor)
	if c.right != nil && c.right.isNotNil() {
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

func (c Comparison) getConditionType() string {
	return "Comparison"
}

func nestedIfCondition(expression Expression) Expression {
	_, isCondition := expression.(Condition)
	if isCondition {
		return NestedExpressionCreate(expression)
	}
	return expression
}
