package cypher_go_dsl

import "fmt"

type FunctionArgumentList struct {
	expressions []Visitable
	key         string
	notNil      bool
}

func (v FunctionArgumentList) isNotNil() bool {
	return v.notNil
}

func (v FunctionArgumentList) getKey() string {
	return v.key
}

func (v FunctionArgumentList) PrepareVisit(child Visitable) Visitable {
	expression, isExpression := child.(Expression)
	if !isExpression {
		return child
	}
	return NameOrExpression(expression)
}

func (v FunctionArgumentList) accept(visitor *CypherRenderer) {
	v.key = fmt.Sprint(&v)
	(*visitor).enter(v)
	for _, expression := range v.expressions {
		v.PrepareVisit(expression).accept(visitor)
	}
	(*visitor).leave(v)
}

func (v FunctionArgumentList) enter(renderer *CypherRenderer) {
}

func (v FunctionArgumentList) leave(renderer *CypherRenderer) {
}
