package cypher_go_dsl

type Pattern struct {
	PatternElements []PatternElement
}

func (p Pattern) PrepareVisit(visitable Visitable) Visitable {
	return visitable
}

func (p Pattern) Accept(visitor *CypherRenderer) {
	(*visitor).Enter(p)
	for _, pattern := range p.PatternElements  {
		p.PrepareVisit(pattern).Accept(visitor)
	}
	(*visitor).Leave(p)
}

func (p Pattern) GetType() VisitableType {
	return PatternVisitable
}

func (p Pattern) Enter(renderer *CypherRenderer) {
}

func (p Pattern) Leave(renderer *CypherRenderer) {
}




