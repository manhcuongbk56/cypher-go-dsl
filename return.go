package cypher

type Return struct {
	distinct Distinct
	body     ReturnBody
	key      string
	notNil   bool
	err      error
}

func ReturnCreate(distinctInstance Distinct, body ReturnBody) Return {
	if body.GetError() != nil {
		return ReturnError(body.GetError())
	}
	r := Return{
		distinct: distinctInstance,
		body:     body,
		notNil:   true,
	}
	r.key = getAddress(&r)
	return r
}

func ReturnCreate1(distinct bool, returnItems ExpressionList, order Order, skip Skip, limit Limit) Return {
	var distinctInstance Distinct
	if distinct {
		distinctInstance = DISTINCT_INSTANCE
	}
	body := ReturnBodyCreate(returnItems, order, skip, limit)
	return ReturnCreate(distinctInstance, body)
}

func ReturnError(err error) Return {
	return Return{
		err: err,
	}
}

func (r Return) GetError() error {
	return r.err
}

func (r Return) isNotNil() bool {
	return r.notNil
}

func (r Return) getKey() string {
	return r.key
}

func (r Return) accept(visitor *CypherRenderer) {
	visitor.enter(r)
	VisitIfNotNull(r.distinct, visitor)
	r.body.accept(visitor)
	visitor.leave(r)
}

func (r Return) enter(renderer *CypherRenderer) {
	renderer.append("RETURN ")
}

func (r Return) leave(renderer *CypherRenderer) {
}
