package cypher_go_dsl

type Properties struct {
	properties MapExpression
	key        string
	notNil     bool
	err        error
}

func PropertiesCreate(newProperties MapExpression) Properties {
	if newProperties.getError() != nil {
		return PropertiesError(newProperties.getError())
	}
	p := Properties{properties: newProperties, notNil: true}
	p.key = getAddress(&p)
	return p
}

func PropertiesError(err error) Properties {
	return Properties{
		err: err,
	}
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
	(*visitor).enter(p)
	p.properties.accept(visitor)
	(*visitor).leave(p)
}

func (p Properties) enter(renderer *CypherRenderer) {
	renderer.append(" ")
}

func (p Properties) leave(renderer *CypherRenderer) {
}
