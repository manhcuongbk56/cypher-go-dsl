package cypher_go_dsl

import (
	"errors"
	"fmt"
)

type Comparison struct {
	left     Expression
	operator Operator
	right    Expression
	key      string
	notNil   bool
	err      error
}

func (c Comparison) getError() error {
	return c.err
}

func (c Comparison) isNotNil() bool {
	return c.notNil
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
	return Comparison{
		left:     left,
		operator: operator,
		right:    right,
	}
}

func ComparisonCreate1(operator Operator, expression Expression) Comparison {
	switch operator.operatorType {
	case PREFIX:
		return Comparison{
			left:     nil,
			operator: operator,
			right:    expression,
		}
	case POSTFIX:
		return Comparison{
			left:     expression,
			operator: operator,
			right:    nil,
		}
	default:
		return Comparison{}
	}

}

func (c Comparison) accept(visitor *CypherRenderer) {
	c.key = fmt.Sprint(&c)
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
