package cypher_go_dsl

import "fmt"

type MultiPartQuery struct {
	parts     []MultiPartElement
	remainder SinglePartQuery
	key       string
	notNil    bool
	err       error
}

func MultiPartQueryCreate(parts []MultiPartElement, remainder SinglePartQuery) MultiPartQuery {
	m := MultiPartQuery{
		parts:     parts,
		remainder: remainder,
	}
	m.key = getAddress(&m)
	return m
}

func (m MultiPartQuery) getError() error {
	return m.err
}

func (m MultiPartQuery) accept(visitor *CypherRenderer) {
	for _, part := range m.parts {
		part.accept(visitor)
	}
	m.remainder.accept(visitor)
}

func (m MultiPartQuery) enter(renderer *CypherRenderer) {
}

func (m MultiPartQuery) leave(renderer *CypherRenderer) {
}

func (m MultiPartQuery) getKey() string {
	return m.key
}

func (m MultiPartQuery) isNotNil() bool {
	return m.notNil
}

func (m MultiPartQuery) doesReturnElements() bool {
	return m.remainder.doesReturnElements()
}
