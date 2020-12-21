package cypher

import (
	"strings"
)

type Namespace struct {
	content []string
	key     string
	notNil  bool
	err     error
}

func NameSpaceCreate(content []string) Namespace {
	n := Namespace{
		content: content,
		notNil:  true,
	}
	n.key = getAddress(&n)
	return n
}

func (n Namespace) getError() error {
	return n.err
}

func (n Namespace) accept(visitor *CypherRenderer) {
	visitor.enter(n)
	visitor.leave(n)
}

func (n Namespace) enter(renderer *CypherRenderer) {
	renderer.append(n.AsString())
}

func (n Namespace) leave(renderer *CypherRenderer) {
}

func (n Namespace) getKey() string {
	return n.key
}

func (n Namespace) isNotNil() bool {
	return n.notNil
}

func (n Namespace) GetExpressionType() ExpressionType {
	return LITERAL
}

func (n Namespace) GetContent() interface{} {
	return n.content
}

func (n Namespace) AsString() string {
	return strings.Join(n.content[:], ".")
}
