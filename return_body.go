package cypher

type ReturnBody struct {
	returnItems ExpressionList
	order       Order
	skip        Skip
	limit       Limit
	key         string
	notNil      bool
	err         error
}

func ReturnBodyCreate(returnItems ExpressionList, order Order, skip Skip, limit Limit) ReturnBody {
	if returnItems.getError() != nil {
		ReturnBodyError(returnItems.getError())
	}
	if order.getError() != nil {
		ReturnBodyError(order.getError())
	}
	if skip.getError() != nil {
		ReturnBodyError(skip.getError())
	}
	if limit.getError() != nil {
		ReturnBodyError(limit.getError())
	}
	r := ReturnBody{
		returnItems: returnItems,
		order:       order,
		skip:        skip,
		limit:       limit,
		notNil:      true,
	}
	r.key = getAddress(&r)
	return r
}

func ReturnBodyError(err error) ReturnBody {
	return ReturnBody{
		err: err,
	}
}

func (r ReturnBody) getError() error {
	return r.err
}

func (r ReturnBody) isNotNil() bool {
	return r.notNil
}

func (r ReturnBody) enter(renderer *CypherRenderer) {
}

func (r ReturnBody) leave(renderer *CypherRenderer) {
}

func (r ReturnBody) getKey() string {
	return r.key
}

func (r ReturnBody) accept(visitor *CypherRenderer) {
	r.returnItems.accept(visitor)
	VisitIfNotNull(r.order, visitor)
	VisitIfNotNull(r.skip, visitor)
	VisitIfNotNull(r.limit, visitor)
}
