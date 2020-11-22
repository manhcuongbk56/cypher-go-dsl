package cypher_go_dsl

type EntryExpression struct {
	Expression
	Key   string
	Value IsExpression
}

func (e EntryExpression) IsExpression() bool {
	return true
}

func (e EntryExpression) Accept(visitor Visitor) {
	visitor.Enter(e)
	e.Value.Accept(visitor)
	visitor.Leave(e)
}

