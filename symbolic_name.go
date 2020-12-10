package cypher_go_dsl

import "fmt"

type SymbolicName struct {
	value  string
	key    string
	notNil bool
	err    error
}

func SymbolicNameCreate(value string) SymbolicName {
	symbolicName := SymbolicName{
		value:  value,
		notNil: true,
	}
	symbolicName.key = getAddress(&symbolicName)
	return symbolicName
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
	(*visitor).enter(s)
	(*visitor).leave(s)
}

func (s SymbolicName) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(renderer.resolve(s))
}

func (s SymbolicName) leave(renderer *CypherRenderer) {
}
