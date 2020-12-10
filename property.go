package cypher_go_dsl

import "fmt"

type Property struct {
	container Expression
	name      PropertyLookup
	key       string
	notNil    bool
	err       error
}

func (p Property) getError() error {
	return p.err
}

func (p Property) isNotNil() bool {
	return p.notNil
}

func (p Property) accept(visitor *CypherRenderer) {
	visitor.enter(p)
	p.container.accept(visitor)
	p.name.accept(visitor)
	visitor.leave(p)
}

func (p Property) enter(renderer *CypherRenderer) {
	panic("implement me")
}

func (p Property) leave(renderer *CypherRenderer) {
	panic("implement me")
}

func (p Property) getKey() string {
	return p.key
}

func (p Property) GetExpressionType() ExpressionType {
	panic("implement me")
}

func (p Property) getName() PropertyLookup {
	return p.name
}
