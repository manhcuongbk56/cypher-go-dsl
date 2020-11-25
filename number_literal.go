package cypher_go_dsl

import "strconv"

type NumberLiteral struct {
	ExpressionStruct
	content int
}

func (n NumberLiteral) IsExpression() bool {
	return true
}

func (n NumberLiteral) GetContent() interface{} {
	return n.content
}

func (n NumberLiteral) AsString() string {
	return strconv.Itoa(n.content)
}

func (n NumberLiteral) Accept(visitor *CypherRenderer) {
	(*visitor).Enter(n)
	(*visitor).Leave(n)
}

func (n NumberLiteral) GetType() VisitableType {
	return LiteralVisitable
}

func (n NumberLiteral) Enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(n.AsString())}

func (n NumberLiteral) Leave(renderer *CypherRenderer) {
}

