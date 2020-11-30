package cypher_go_dsl

import "fmt"

type Pattern struct {
	patternElements []PatternElement
	key             string
}

func (p Pattern) getKey() string {
	return p.key
}

func (p Pattern) PrepareVisit(visitable Visitable) Visitable {
	return visitable
}

func (p Pattern) accept(visitor *CypherRenderer) {
	p.key = fmt.Sprint(&p)
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
