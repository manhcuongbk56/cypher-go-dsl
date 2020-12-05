package cypher_go_dsl

import (
	"fmt"
	"strings"
)

type Namespace struct {
	content []string
	key     string
	notNil  bool
}

func NameSpaceCreate(content []string) Namespace {
	return Namespace{
		content: content,
		notNil:  true,
	}
}

func (n Namespace) accept(visitor *CypherRenderer) {
	n.key = fmt.Sprint(&n)
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
