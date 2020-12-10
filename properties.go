package cypher_go_dsl

import "fmt"

type Properties struct {
	properties MapExpression
	key        string
	notNil     bool
	err error
}

func (p Properties) isNotNil() bool {
	return p.notNil
}

func (p Properties) getKey() string {
	return p.key
}

func (p Properties) getError() error {
	return p.err
}

func (p Properties) accept(visitor *CypherRenderer) {
	p.key = fmt.Sprint(&p)
	(*visitor).enter(p)
	p.properties.accept(visitor)
	(*visitor).leave(p)
}

func (p Properties) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(" ")
}

func (p Properties) leave(renderer *CypherRenderer) {
}
