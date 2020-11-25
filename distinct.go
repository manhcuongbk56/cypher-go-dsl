package cypher_go_dsl

type Distinct struct {
	IsDistinct bool
}

func (d Distinct) Accept(visitor *CypherRenderer) {
	(*visitor).Enter(&d)
	(*visitor).Leave(&d)
}

func (d Distinct) GetType() VisitableType {
	return DistinctVisitable
}

func (d Distinct) Enter(renderer *CypherRenderer) {
	panic("implement me")
}

func (d Distinct) Leave(renderer *CypherRenderer) {
	panic("implement me")
}



