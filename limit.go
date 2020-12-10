package cypher_go_dsl

import "fmt"

type Limit struct {
	limitAmount NumberLiteral
	key         string
	notNil      bool
	err error
}

func (l Limit) getError() error {
	return l.err
}

func (l Limit) isNotNil() bool {
	return l.notNil
}

func (l Limit) getKey() string {
	return l.key
}

func (l Limit) accept(visitor *CypherRenderer) {
	l.key = fmt.Sprint(&l)
	(*visitor).enter(l)
	l.limitAmount.accept(visitor)
	(*visitor).leave(l)
}

func CreateLimit(number int) Limit {
	if number == 0 {
		return Limit{}
	}
	literal := NumberLiteral{
		content: number,
	}
	return Limit{limitAmount: literal}
}

func (l Limit) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(" LIMIT ")
}

func (l Limit) leave(renderer *CypherRenderer) {
}
