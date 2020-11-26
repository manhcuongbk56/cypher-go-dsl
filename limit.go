package cypher_go_dsl

import "fmt"

type Limit struct {
	limitAmount NumberLiteral
	key         string
}

func (l Limit) getKey() string {
	return l.key
}

func (l Limit) accept(visitor *CypherRenderer) {
	l.key = fmt.Sprint(&l)
	(*visitor).Enter(l)
	l.limitAmount.accept(visitor)
	(*visitor).Leave(l)
}

func CreateLimit(number int) *Limit {
	if number == 0 {
		return nil
	}
	literal := NumberLiteral{
		content: number,
	}
	return &Limit{limitAmount: literal}
}

func (l Limit) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(" LIMIT ")
}

func (l Limit) leave(renderer *CypherRenderer) {
}
