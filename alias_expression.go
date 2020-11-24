package cypher_go_dsl

type AliasedExpression struct {
	delegate Expression
	alias    string
}

func (aliased AliasedExpression) As(newAlias string) AliasedExpression {
	return AliasedExpression{delegate: aliased.delegate,
		alias: newAlias}
}

func (aliased AliasedExpression) Accept(visitor Visitor) {
	visitor.Enter(aliased)
	NameOrExpression(aliased.delegate).Accept(visitor)
	visitor.Leave(aliased)
}

func (aliased AliasedExpression) GetType() VisitableType {
	panic("implement me")
}
