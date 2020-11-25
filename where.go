package cypher_go_dsl

type Where struct {
	condition Condition
}

func (w Where) Accept(visitor *CypherRenderer) {
	(*visitor).Enter(w)
	w.Accept(visitor)
	(*visitor).Leave(w)
}

func (w Where) GetType() VisitableType {
	return WhereVisitable
}

func (w Where) Enter(renderer *CypherRenderer) {
	panic("implement me")
}

func (w Where) Leave(renderer *CypherRenderer) {
	panic("implement me")
}



