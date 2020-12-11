package cypher_go_dsl

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
	if !operator.isUnary() {
		return Comparison{
			err: errors.New("comparison: operator must be unary"),
		}
	}
	if left == nil || !left.isNotNil() {
		return Comparison{
			err: errors.New("comparison: left expression must be not nil"),
		}
	}
	if right == nil || !left.isNotNil() {
		return Comparison{
			err: errors.New("comparison: right expression must be not nil"),
		}
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
	var comparision Comparison
	switch operator.operatorType {
	case PREFIX:
		comparision = Comparison{
			left:     nil,
			operator: operator,
			right:    expression,
		}
	case POSTFIX:
		comparision = Comparison{
			left:     expression,
			operator: operator,
			right:    nil,
		}
	default:
		return Comparison{}
	}
	comparision.key = getAddress(&comparision)
	return comparision
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
