package cypher_go_dsl

type With struct {
	distinct Distinct
	body     ReturnBody
	where    Where
	key      string
	notNil   bool
	err      error
}

func WithCreate(distinct bool, returnItems ExpressionList, order Order, skip Skip, limit Limit, where Where) With {
	if returnItems.getError() != nil {
		WithError(returnItems.getError())
	}
	if order.getError() != nil {
		WithError(order.getError())
	}
	if skip.getError() != nil {
		WithError(skip.getError())
	}
	if limit.getError() != nil {
		WithError(limit.getError())
	}
	if where.getError() != nil {
		WithError(where.getError())
	}
	distinctInstance := Distinct{}
	if distinct {
		distinctInstance = DISTINCT_INSTANCE
	}
	with := With{
		distinct: distinctInstance,
		body:     ReturnBodyCreate(returnItems, order, skip, limit),
		where:    where,
		notNil:   true,
	}
	with.key = getAddress(&with)
	return with
}

func WithError(err error) With {
	return With{
		err: err,
	}
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
