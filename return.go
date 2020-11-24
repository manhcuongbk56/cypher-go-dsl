package cypher_go_dsl

type Return struct {
	distinct *Distinct
	body *ReturnBody
}

func (r Return) Accept(visitor Visitor) {
	visitor.Enter(r)
	VisitIfNotNull(r.distinct, visitor)
	r.body.Accept(visitor)
	visitor.Leave(r)
}

func ReturnByMultiVariable(distinct bool, returnItems ExpressionList, order *Order, skip *Skip, limit *Limit) *Return {
	var distinctInstance *Distinct
	if distinct {
		distinctInstance = &Distinct{}
	}
	body := ReturnBody{
		returnItems: returnItems,
		order:       order,
		skip:        skip,
		limit:       limit,
	}
	return &Return{
		distinct: distinctInstance,
		body:     &body,
	}
}

