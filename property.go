package cypher_go_dsl

import "fmt"

type Property struct {
	container Expression
	name PropertyLookup
	key string
}

func (p Property) accept(visitor *CypherRenderer) {
	p.key = fmt.Sprint(&p)
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
