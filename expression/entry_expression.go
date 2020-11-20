package expression

import (
	v "cypher-go-dsl/visitable"
)

type EntryExpression struct {
	Expression
	Key string
	Value HasExpression
}

func (e EntryExpression) GetExpression() Expression {
	return Expression{}
}

func (e EntryExpression) Accept(visitor v.Visitor) {
	visitor.Enter(e)
	e.Value.Accept(visitor)
	visitor.Leave(e)
}

