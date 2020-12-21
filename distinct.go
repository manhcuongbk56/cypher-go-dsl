package cypher

type Distinct struct {
	IsDistinct bool
	key        string
	notNil     bool
	err        error
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
	(*visitor).enter(d)
	(*visitor).leave(d)
}

func (d Distinct) enter(renderer *CypherRenderer) {
	renderer.append("DISTINCT ")
}

func (d Distinct) leave(renderer *CypherRenderer) {
	panic("implement me")
}
