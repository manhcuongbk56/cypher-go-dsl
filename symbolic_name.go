package cypher

import (
	"errors"
)

type SymbolicName struct {
	ExpressionContainer
	value  string
	key    string
	notNil bool
	err    error
}

func SymbolicNameCreate(value string) SymbolicName {
	if value == "" {
		return SymbolicNameError(errors.New("name must be not empty"))
	}
	symbolicName := SymbolicName{
		value:  value,
		notNil: true,
	}
	symbolicName.key = getAddress(&symbolicName)
	symbolicName.ExpressionContainer = ExpressionWrap(symbolicName)
	return symbolicName
}

func SymbolicNameError(err error) SymbolicName {
	return SymbolicName{
		err: err,
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
	(*visitor).enter(s)
	(*visitor).leave(s)
}

func (s SymbolicName) enter(renderer *CypherRenderer) {
	renderer.append(renderer.resolve(s))
}

func (s SymbolicName) leave(renderer *CypherRenderer) {
}
