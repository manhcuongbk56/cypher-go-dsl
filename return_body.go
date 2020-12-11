package cypher_go_dsl

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

func (r ReturnBody) getError() error {
	return r.err
}

func (r ReturnBody) isNotNil() bool {
	return r.notNil
}

func (r ReturnBody) enter(renderer *CypherRenderer) {
	panic("implement me")
}

func (r ReturnBody) leave(renderer *CypherRenderer) {
	panic("implement me")
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
