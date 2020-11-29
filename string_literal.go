package cypher_go_dsl

import (
	"fmt"
)

type StringLiteral struct {
	ExpressionStruct
	content string
	key     string
}

func escapeStringLiteral(value string) string {
	return "'" + value + "'"
}

func (s StringLiteral) getKey() string {
	return s.key
}

func (s StringLiteral) IsExpression() bool {
	return true
}

func (s StringLiteral) GetContent() interface{} {
	return s.content
}

func (s StringLiteral) AsString() string {
	return s.content
}

func (s StringLiteral) accept(visitor *CypherRenderer) {
	s.key = fmt.Sprint(&s)
	(*visitor).enter(s)
	(*visitor).Leave(s)
}

func (s StringLiteral) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(escapeStringLiteral(s.AsString()))
}

func (s StringLiteral) leave(renderer *CypherRenderer) {
}
