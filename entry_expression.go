package cypher_go_dsl

import "fmt"

type EntryExpression struct {
	Key    string
	Value  Expression
	key    string
	notNil bool
	err    error
}

func EntryExpressionCreate(key string, value Expression) EntryExpression {
	e := EntryExpression{
		Value:  value,
		Key:    key,
		notNil: true,
	}
	e.key = getAddress(&e)
	return e
}

func (e EntryExpression) getError() error {
	return e.err
}

func (e EntryExpression) isNotNil() bool {
	return e.notNil
}

func (e EntryExpression) getKey() string {
	return e.key
}

func (e EntryExpression) GetExpressionType() ExpressionType {
	return EXPRESSION
}

func (e EntryExpression) accept(visitor *CypherRenderer) {
	(*visitor).enter(e)
	e.Value.accept(visitor)
	(*visitor).leave(e)
}

func (e EntryExpression) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(escapeName(e.Key))
	renderer.builder.WriteString(": ")
}

func (e EntryExpression) leave(renderer *CypherRenderer) {
}
