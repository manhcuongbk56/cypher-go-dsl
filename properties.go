package cypher_go_dsl

import "fmt"

type Properties struct {
	properties MapExpression
	key        string
}

func (p Properties) getKey() string {
	return p.key
}

func (p Properties) accept(visitor *CypherRenderer) {
	p.key = fmt.Sprint(&p)
	(*visitor).Enter(p)
	p.properties.accept(visitor)
	(*visitor).Leave(p)
}

func (p Properties) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(" ")
}

func (p Properties) leave(renderer *CypherRenderer) {
}
