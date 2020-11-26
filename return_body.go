package cypher_go_dsl

import "fmt"

type ReturnBody struct {
	returnItems ExpressionList
	order       *Order
	skip        *Skip
	limit       *Limit
	key         string
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
	r.key = fmt.Sprint(&r)
	r.returnItems.accept(visitor)
	VisitIfNotNull(r.order, visitor)
	VisitIfNotNull(r.skip, visitor)
	VisitIfNotNull(r.limit, visitor)
}
