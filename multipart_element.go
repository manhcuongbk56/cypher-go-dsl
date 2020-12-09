package cypher_go_dsl

import "fmt"

type MultiPartElement struct {
	precedingClauses []Visitable
	with             With
	key              string
	notNil           bool
}

func MultiPartElementCreate(precedingClauses []Visitable, with With) MultiPartElement {
	var clauses []Visitable
	if precedingClauses == nil || len(precedingClauses) == 0 {
		clauses = make([]Visitable, 0)
	} else {
		clauses = make([]Visitable, 0)
		clauses = append(clauses, precedingClauses...)
	}
	return MultiPartElement{
		precedingClauses: clauses,
		with:             with,
		notNil:           true,
	}
}

func (m MultiPartElement) accept(visitor *CypherRenderer) {
	m.key = fmt.Sprint(&m)
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
