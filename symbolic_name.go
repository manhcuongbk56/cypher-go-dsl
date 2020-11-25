package cypher_go_dsl

type SymbolicName struct {
	Value string
}

func (s SymbolicName) IsExpression() bool {
	return true
}

func (s SymbolicName) Accept(visitor *CypherRenderer) {
	(*visitor).Enter(s)
	(*visitor).Leave(s)
}

func (s SymbolicName) GetType() VisitableType {
	return SymbolicNameVisitable
}

func (s SymbolicName) Enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(renderer.resolve(s))}

func (s SymbolicName) Leave(renderer *CypherRenderer) {
}
