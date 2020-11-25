package cypher_go_dsl

type StringLiteral struct {
	ExpressionStruct
	content string
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

func (s StringLiteral) Accept(visitor *CypherRenderer) {
	(*visitor).Enter(s)
	(*visitor).Leave(s)
}

func (s StringLiteral) GetType() VisitableType {
	return StringLiteralVisitable
}

func (s StringLiteral) Enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(s.AsString())}

func (s StringLiteral) Leave(renderer *CypherRenderer) {
}

