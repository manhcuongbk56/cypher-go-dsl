package cypher_go_dsl

import "fmt"

type Set struct {
	setItems ExpressionList
	key      string
	notNil   bool
}

func SetCreate(setItems ExpressionList) Set {
	return Set{
		setItems: setItems,
		notNil:   true,
	}
}

func (s Set) accept(visitor *CypherRenderer) {
	s.key = fmt.Sprint(&s)
	visitor.enter(s)
	s.setItems.accept(visitor)
	visitor.leave(s)
}

func (s Set) enter(renderer *CypherRenderer) {
}

func (s Set) leave(renderer *CypherRenderer) {
}

func (s Set) getKey() string {
	return s.key
}

func (s Set) isNotNil() bool {
	return s.notNil
}
