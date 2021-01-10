package cypher

type with struct {
	distinct Distinct
	body     ReturnBody
	where    Where
	key      string
	notNil   bool
	err      error
}

func withCreate(distinct bool, returnItems ExpressionList, order Order, skip Skip, limit Limit, where Where) with {
	if returnItems.GetError() != nil {
		withError(returnItems.GetError())
	}
	if order.GetError() != nil {
		withError(order.GetError())
	}
	if skip.GetError() != nil {
		withError(skip.GetError())
	}
	if limit.GetError() != nil {
		withError(limit.GetError())
	}
	if where.GetError() != nil {
		withError(where.GetError())
	}
	distinctInstance := Distinct{}
	if distinct {
		distinctInstance = DISTINCT_INSTANCE
	}
	with := with{
		distinct: distinctInstance,
		body:     ReturnBodyCreate(returnItems, order, skip, limit),
		where:    where,
		notNil:   true,
	}
	with.key = getAddress(&with)
	return with
}

func withError(err error) with {
	return with{
		err: err,
	}
}

func (with with) GetError() error {
	return with.err
}

func (with with) accept(visitor *CypherRenderer) {
	visitor.enter(with)
	VisitIfNotNull(with.distinct, visitor)
	with.body.accept(visitor)
	VisitIfNotNull(with.where, visitor)
	visitor.leave(with)
}

func (with with) enter(renderer *CypherRenderer) {
	renderer.append("WITH ")
}

func (with with) leave(renderer *CypherRenderer) {
	renderer.append(" ")
}

func (with with) getKey() string {
	return with.key
}

func (with with) isNotNil() bool {
	return with.notNil
}
