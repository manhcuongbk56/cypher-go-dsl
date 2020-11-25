package cypher_go_dsl

type Properties struct {
	properties MapExpression
}

func (p Properties) Accept(visitor *CypherRenderer) {
	(*visitor).Enter(p)
	p.properties.Accept(visitor)
	(*visitor).Leave(p)
}

func (p Properties) GetType() VisitableType {
	return PropertiesVisitable
}

func (p Properties) Enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(" ")
}

func (p Properties) Leave(renderer *CypherRenderer) {
}

