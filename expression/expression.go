package expression

import (
	v "cypher-go-dsl/visitable"
)

type HasExpression interface {
	GetExpression() Expression
	Accept(visitor v.Visitor)
}

type Expression struct {
}

type Condition struct {
	expression Expression
}

//type ICondition interface {
//	IExpression
//}

func (lhs Expression) IsEqualTo(rhs Expression) Comparison {
	return Comparison{left: lhs, operator: EQUALITY, right: rhs}
}

func (lhs Expression) Accept(visitor v.Visitor) {
	panic("implement me")
}

func (lhs Condition) And(rhs Condition)  {

}


type Operator string

const (
	EQUALITY = "equality"
)

type Comparison struct {
	Expression
	left Expression
	operator Operator
	right Expression
}

func (c Comparison) GetExpression() Expression {
	return Expression{}
}

func (c Comparison) Accept(visitor v.Visitor) {
	panic("implement me")
}



