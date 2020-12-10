package cypher_go_dsl

import "fmt"

type SymbolicName struct {
	value  string
	key    string
	notNil bool
	err error
}

func SymbolicNameCreate(value string) SymbolicName {
	return SymbolicName{
		value:  value,
		notNil: true,
	}
}

func (s SymbolicName) getError() error {
	return s.err
}

func (s SymbolicName) isNotNil() bool {
	return s.notNil
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
