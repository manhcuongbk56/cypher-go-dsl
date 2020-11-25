package cypher_go_dsl

type Return struct {
	distinct *Distinct
	body *ReturnBody
}

func (r Return) Accept(visitor *CypherRenderer) {
	(*visitor).Enter(r)
	VisitIfNotNull(r.distinct, visitor)
	r.body.Accept(visitor)
	(*visitor).Leave(r)
}

func (r Return) GetType() VisitableType {
	return ReturnVisitable
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

func (r Return) Enter(renderer *CypherRenderer) {
	renderer.builder.WriteString("RETURN  ")}

func (r Return) Leave(renderer *CypherRenderer) {
}

