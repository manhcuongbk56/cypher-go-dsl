package cypher_go_dsl

type EntryExpression struct {
	ExpressionStruct
	Key   string
	Value Expression
}

func (e EntryExpression) IsExpression() bool {
	return true
}

func (e EntryExpression) Accept(visitor Visitor) {
	visitor.Enter(e)
	e.Value.Accept(visitor)
	visitor.Leave(e)
}

func (e EntryExpression) GetType() VisitableType {
	return EntryExpressionVisitable
}

