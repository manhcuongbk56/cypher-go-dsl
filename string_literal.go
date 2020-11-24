package cypher_go_dsl

type StringLiteral struct {
	ExpressionStruct
	content string
}

func (n StringLiteral) IsExpression() bool {
	return true
}

func (n StringLiteral) GetContent() interface{} {
	return n.content
}

func (n StringLiteral) AsString() string {
	return n.content
}

func (n StringLiteral) Accept(visitor Visitor) {
	visitor.Enter(n)
	visitor.Leave(n)
}

func (n StringLiteral) GetType() VisitableType {
	return StringLiteralVisitable
}

