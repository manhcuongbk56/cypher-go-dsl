package cypher

import "errors"

type YieldItems struct {
	expressions []Expression
	key         string
	notNil      bool
	err         error
}

func YieldItemsCreate(expressions []Expression) YieldItems {
	if expressions == nil || len(expressions) == 0 {
		return YieldItemsError(errors.New("can not yield an empty list of items"))
	}
	for _, expression := range expressions {
		if expression.getError() != nil {
			return YieldItemsError(expression.getError())
		}
	}
	yieldItem := YieldItems{
		expressions: expressions,
		notNil:      true,
	}
	yieldItem.key = getAddress(&yieldItem)
	return yieldItem
}

func YieldItemsError(err error) YieldItems {
	return YieldItems{
		err: nil,
	}
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
