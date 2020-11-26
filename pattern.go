package cypher_go_dsl

import "fmt"

type Pattern struct {
	PatternElements []PatternElement
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
	(*visitor).Enter(p)
	for _, pattern := range p.PatternElements {
		p.PrepareVisit(pattern).accept(visitor)
	}
	(*visitor).Leave(p)
}

func (p Pattern) enter(renderer *CypherRenderer) {
}

func (p Pattern) leave(renderer *CypherRenderer) {
}
