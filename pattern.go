package cypher

type Pattern struct {
	patternElements []PatternElement
	key             string
	notNil          bool
	err             error
}

func PatternCreate(patternElements []PatternElement) Pattern {
	for _, element := range patternElements {
		if element != nil && element.GetError() != nil {
			return Pattern{err: element.GetError()}
		}
	}
	p := Pattern{
		patternElements: patternElements,
		notNil:          true,
	}
	p.key = getAddress(&p)
	return p
}

func (p Pattern) GetError() error {
	return p.err
}

func (p Pattern) isNotNil() bool {
	return p.notNil
}

func (p Pattern) getKey() string {
	return p.key
}

func (p Pattern) PrepareVisit(visitable Visitable) Visitable {
	return visitable
}

func (p Pattern) accept(visitor *CypherRenderer) {
	(*visitor).enter(p)
	for _, pattern := range p.patternElements {
		p.PrepareVisit(pattern).accept(visitor)
	}
	(*visitor).leave(p)
}

func (p Pattern) enter(renderer *CypherRenderer) {
}

func (p Pattern) leave(renderer *CypherRenderer) {
}
