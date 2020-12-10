package cypher_go_dsl

import "fmt"

type Distinct struct {
	IsDistinct bool
	key        string
	notNil     bool
	err error
}

var DISTINCT_INSTANCE = Distinct{
	IsDistinct: true,
	key:        "DISTINCT_INSTANCE",
}

func (d Distinct) getError() error {
	return d.err
}

func (d Distinct) isNotNil() bool {
	return d.notNil
}

func (d Distinct) getKey() string {
	return d.key
}

func (d Distinct) accept(visitor *CypherRenderer) {
	d.key = fmt.Sprint(&d)
	(*visitor).enter(d)
	(*visitor).leave(d)
}

func (d Distinct) enter(renderer *CypherRenderer) {
	panic("implement me")
}

func (d Distinct) leave(renderer *CypherRenderer) {
	panic("implement me")
}
