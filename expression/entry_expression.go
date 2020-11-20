package expression

import (
	v "../visitable"
)

type EntryExpression struct {
	Key string
	Value Expression
}

func (e EntryExpression) Accept(visitor v.Visitor) {
	visitor.Enter(e)
	e.Value.Accept(visitor)
	visitor.Leave(e)
}

