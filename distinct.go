package cypher_go_dsl

type Distinct struct {
	IsDistinct bool
}

func (d Distinct) Accept(visitor Visitor) {
	visitor.Enter(d)
	visitor.Leave(d)
}

