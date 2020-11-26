package cypher_go_dsl

import "fmt"

type EntryExpression struct {
	ExpressionStruct
	Key   string
	Value Expression
	key   string
}

func (e EntryExpression) getKey() string {
	return e.key
}

func (e EntryExpression) IsExpression() bool {
	return true
}

func (e EntryExpression) accept(visitor *CypherRenderer) {
	e.key = fmt.Sprint(&e)
	(*visitor).Enter(e)
	e.Value.accept(visitor)
	(*visitor).Leave(e)
}

func (e EntryExpression) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(escapeName(e.Key))
	renderer.builder.WriteString(": ")
}

func (e EntryExpression) leave(renderer *CypherRenderer) {
}
