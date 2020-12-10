package cypher_go_dsl

import "fmt"

type MultiPartElement struct {
	precedingClauses []Visitable
	with             With
	key              string
	notNil           bool
	err              error
}

func MultiPartElementCreate(precedingClauses []Visitable, with With) MultiPartElement {
	var clauses []Visitable
	if precedingClauses == nil || len(precedingClauses) == 0 {
		clauses = make([]Visitable, 0)
	} else {
		clauses = make([]Visitable, 0)
		clauses = append(clauses, precedingClauses...)
	}
	m := MultiPartElement{
		precedingClauses: clauses,
		with:             with,
		notNil:           true,
	}
	m.key = getAddress(&m)
	return m
}

func (m MultiPartElement) getError() error {
	return m.err
}

func (m MultiPartElement) accept(visitor *CypherRenderer) {
	visitor.enter(m)
	for _, clause := range m.precedingClauses {
		clause.accept(visitor)
	}
	m.with.accept(visitor)
	visitor.leave(m)
}

func (m MultiPartElement) enter(renderer *CypherRenderer) {
}

func (m MultiPartElement) leave(renderer *CypherRenderer) {
}

func (m MultiPartElement) getKey() string {
	return m.key
}

func (m MultiPartElement) isNotNil() bool {
	return m.notNil
}
