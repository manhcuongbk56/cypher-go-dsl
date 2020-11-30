package cypher_go_dsl

import "fmt"

type SymbolicName struct {
	value string
	key   string
}

func (s SymbolicName) getKey() string {
	return s.key
}

func (s SymbolicName) GetExpressionType() ExpressionType {
	return EXPRESSION
}

func (s SymbolicName) accept(visitor *CypherRenderer) {
	s.key = fmt.Sprint(&s)
	(*visitor).enter(s)
	(*visitor).leave(s)
}

func (s SymbolicName) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(renderer.resolve(s))
}

func (s SymbolicName) leave(renderer *CypherRenderer) {
}
