package cypher_go_dsl

import "fmt"

type With struct {
	distinct Distinct
	body     ReturnBody
	where    Where
	key      string
	notNil   bool
	err      error
}

func WithCreate(distinct bool, returnItems ExpressionList, order Order, skip Skip, limit Limit, where Where) With {
	with := With{
		distinct: DISTINCT_INSTANCE,
		body:     ReturnBodyCreate(returnItems, order, skip, limit),
		notNil:   true,
	}
	with.key = getAddress(&with)
	return with
}

func (with With) getError() error {
	return with.err
}

func (with With) accept(visitor *CypherRenderer) {
	visitor.enter(with)
	VisitIfNotNull(with.distinct, visitor)
	with.body.accept(visitor)
	VisitIfNotNull(with.where, visitor)
	visitor.leave(with)
}

func (with With) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString("WITH ")
}

func (with With) leave(renderer *CypherRenderer) {
	renderer.builder.WriteString(" ")
}

func (with With) getKey() string {
	return with.key
}

func (with With) isNotNil() bool {
	return with.notNil
}
