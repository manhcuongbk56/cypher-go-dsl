package cypher_go_dsl

import "fmt"

type Arguments struct {
	expressions []Expression
	key         string
	notNil      bool
}

func ArgumentsCreate(expression []Expression) Arguments {
	return Arguments{
		expressions: expression,
		notNil:      true,
	}
}

func (a Arguments) isNotNil() bool {
	return a.notNil
}

func (a Arguments) getKey() string {
	return a.key
}

func (a Arguments) PrepareVisit(child Visitable) Visitable {
	expression, isExpression := child.(Expression)
	if !isExpression {
		panic("Can not prepare un expression type in expression list")
	}
	return NameOrExpression(expression)
}

func (a Arguments) accept(visitor *CypherRenderer) {
	a.key = fmt.Sprint(&a)
	(*visitor).enter(a)
	for _, expression := range a.expressions {
		a.PrepareVisit(expression).accept(visitor)
	}
	(*visitor).leave(a)
}

func (a Arguments) enter(renderer *CypherRenderer) {
}

func (a Arguments) leave(renderer *CypherRenderer) {
}
