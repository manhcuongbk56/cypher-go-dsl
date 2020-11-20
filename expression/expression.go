package expression

import (
	v  "../visitable"
)

type IExpression interface {
	IsEqualTo(rhs IExpression) Comparison
}

type Expression struct {
}

type Condition struct {
	IExpression
}

type ICondition interface {
	IExpression
}

func (lhs Expression) IsEqualTo(rhs IExpression) Comparison {
	panic("implement me")
}

func (lhs Expression) Accept(visitor v.Visitor) {
	panic("implement me")
}

func (lhs Condition) And(rhs Condition)  {

}

func isEqualTo(lhs Expression, rhs Expression) Comparison {
return Comparison{
	Left: lhs,
	Operator: EQUALITY,
	Right: rhs,
}
}

