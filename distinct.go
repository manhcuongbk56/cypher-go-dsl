package cypher_go_dsl

import "fmt"

type Distinct struct {
	IsDistinct bool
	key        string
}

func (d Distinct) getKey() string {
	return d.key
}

func (d Distinct) accept(visitor *CypherRenderer) {
	d.key = fmt.Sprint(&d)
	(*visitor).enter(&d)
	(*visitor).leave(&d)
}

func (d Distinct) enter(renderer *CypherRenderer) {
	panic("implement me")
}

func (d Distinct) leave(renderer *CypherRenderer) {
	panic("implement me")
}
