package cypher_go_dsl

import "fmt"

type Set struct {
	setItems ExpressionList
	key      string
	notNil   bool
	err error
}

func (s Set) isUpdatingClause() bool {
	return true
}

func SetCreate(setItems ExpressionList) Set {
	return Set{
		setItems: setItems,
		notNil:   true,
	}
}

func (s Set) getError() error {
	return s.err
}

func (s Set) accept(visitor *CypherRenderer) {
	s.key = fmt.Sprint(&s)
	visitor.enter(s)
	s.setItems.accept(visitor)
	visitor.leave(s)
}

func (s Set) enter(renderer *CypherRenderer) {
	renderer.append("SET ")
}

func (s Set) leave(renderer *CypherRenderer) {
	renderer.append(" ")
}

func (s Set) getKey() string {
	return s.key
}

func (s Set) isNotNil() bool {
	return s.notNil
}
