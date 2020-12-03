package cypher_go_dsl

import (
	"fmt"
)

type StringLiteral struct {
	content string
	key     string
	notNil  bool
}

func StringLiteralCreate(content string) StringLiteral {
	return StringLiteral{
		content: content,
		notNil:  true,
	}
}

func (s StringLiteral) isNotNil() bool {
	return s.notNil
}

func escapeStringLiteral(value string) string {
	return "'" + value + "'"
}

func (s StringLiteral) getKey() string {
	return s.key
}

func (s StringLiteral) GetExpressionType() ExpressionType {
	return EXPRESSION
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
	(*visitor).leave(s)
}

func (s StringLiteral) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(escapeStringLiteral(s.AsString()))
}

func (s StringLiteral) leave(renderer *CypherRenderer) {
}
