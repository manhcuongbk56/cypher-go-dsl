package cypher_go_dsl

import (
	"strconv"
)

type NumberLiteral struct {
	content int
	key     string
	notNil  bool
	err     error
}

func NumberLiteralCreate(content int) NumberLiteral {
	n := NumberLiteral{
		content: content,
		notNil:  true,
	}
	n.key = getAddress(&n)
	return n
}

func (n NumberLiteral) getError() error {
	return n.err
}

func (n NumberLiteral) isNotNil() bool {
	return n.notNil
}

func (n NumberLiteral) getKey() string {
	return n.key
}

func (n NumberLiteral) GetExpressionType() ExpressionType {
	return EXPRESSION
}

func (n NumberLiteral) GetContent() interface{} {
	return n.content
}

func (n NumberLiteral) AsString() string {
	return strconv.Itoa(n.content)
}

func (n NumberLiteral) accept(visitor *CypherRenderer) {
	(*visitor).enter(n)
	(*visitor).leave(n)
}

func (n NumberLiteral) enter(renderer *CypherRenderer) {
	renderer.append(n.AsString())
}

func (n NumberLiteral) leave(renderer *CypherRenderer) {
}
