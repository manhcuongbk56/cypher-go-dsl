package cypher_go_dsl

type SymbolicName struct {
	Value string
}

func (s SymbolicName) IsExpression() bool {
	return true
}

func (s SymbolicName) Accept(visitor Visitor) {
	visitor.Enter(s)
	visitor.Leave(s)
}

func (s SymbolicName) GetType() VisitableType {
	return SymbolicNameVisitable
}
