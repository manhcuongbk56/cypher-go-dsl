package cypher_go_dsl

import "fmt"

type YieldItems struct {
	expressions []Expression
	key         string
	notNil      bool
	err         error
}

func YieldItemsCreate(expression []Expression) YieldItems {
	yieldItem := YieldItems{
		expressions: expression,
		notNil:      true,
	}
	yieldItem.key = getAddress(&yieldItem)
	return yieldItem
}

func (e YieldItems) getError() error {
	return e.err
}

func (e YieldItems) isNotNil() bool {
	return e.notNil
}

func (e YieldItems) getKey() string {
	return e.key
}

func (e YieldItems) PrepareVisit(child Visitable) Visitable {
	return child
}

func (e YieldItems) accept(visitor *CypherRenderer) {
	(*visitor).enter(e)
	for _, expression := range e.expressions {
		e.PrepareVisit(expression).accept(visitor)
	}
	(*visitor).leave(e)
}

func (e YieldItems) enter(renderer *CypherRenderer) {
	renderer.append(" YIELD ")
}

func (e YieldItems) leave(renderer *CypherRenderer) {
}

func yieldAllOf(c ...Expression) YieldItems {
	return YieldItemsCreate(c)
}
