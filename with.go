package cypher

type With struct {
	distinct Distinct
	body     ReturnBody
	where    Where
	key      string
	notNil   bool
	err      error
}

func WithCreate(distinct bool, returnItems ExpressionList, order Order, skip Skip, limit Limit, where Where) With {
	if returnItems.GetError() != nil {
		WithError(returnItems.GetError())
	}
	if order.GetError() != nil {
		WithError(order.GetError())
	}
	if skip.GetError() != nil {
		WithError(skip.GetError())
	}
	if limit.GetError() != nil {
		WithError(limit.GetError())
	}
	if where.GetError() != nil {
		WithError(where.GetError())
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

func (with With) GetError() error {
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
	renderer.append("WITH ")
}

func (with With) leave(renderer *CypherRenderer) {
	renderer.append(" ")
}

func (with With) getKey() string {
	return with.key
}

func (with With) isNotNil() bool {
	return with.notNil
}
