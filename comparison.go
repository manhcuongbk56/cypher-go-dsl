package cypher_go_dsl

import (
	"errors"
	"fmt"
)

type Comparison struct {
	left     *Expression
	operator *Operator
	right    *Expression
	key string
}

func NewComparison(left *Expression, operator *Operator, right *Expression) (Comparison, error) {
	if left == nil {
		return Comparison{}, errors.New("left expression must be not nil")
	}
	if operator == nil {
		return Comparison{}, errors.New("operator must be not nil")
	}
	if right == nil {
		return Comparison{}, errors.New("right expression must be not nil")
	}
	return Comparison{
		left:             left,
		operator:         operator,
		right:            right,
	}, nil
}

func (c Comparison) accept(visitor *CypherRenderer) {
	c.key = fmt.Sprint(&c)
	visitor.enter(c)
	if c.left != nil {
		NameOrExpression(c.left).accept(visitor)
	}
	c.operator.Accept(visitor)
	if c.right != nil {
		NameOrExpression(c.right).accept(visitor)
	}
	visitor.Leave(c)
}

func (c Comparison) enter(renderer *CypherRenderer) {
	panic("implement me")
}

func (c Comparison) leave(renderer *CypherRenderer) {
	panic("implement me")
}

func (c Comparison) getKey() string {
	panic("implement me")
}

func (c Comparison) IsExpression() bool {
	return true
}

func NameOrExpression(expression *Expression) Expression {
	named, isNamed := (*expression).(Named)
	if isNamed && named.getSymbolicName() != nil {
		return named.getSymbolicName()
	}
	return *expression
}

