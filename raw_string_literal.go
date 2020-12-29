package cypher

type RawStringLiteral struct {
	ExpressionContainer
	content string
	key     string
	notNil  bool
	err     error
}

func RawStringLiteralCreate(content string) RawStringLiteral {
	RawStringLiteral := RawStringLiteral{
		content: content,
		notNil:  true,
	}
	RawStringLiteral.key = getAddress(&RawStringLiteral)
	RawStringLiteral.ExpressionContainer = ExpressionWrap(RawStringLiteral)
	return RawStringLiteral
}

func RawStringLiteralError(err error) RawStringLiteral {
	return RawStringLiteralError(err)
}

func (s RawStringLiteral) getError() error {
	return s.err
}

func (s RawStringLiteral) isNotNil() bool {
	return s.notNil
}

func (s RawStringLiteral) getKey() string {
	return s.key
}

func (s RawStringLiteral) GetExpressionType() ExpressionType {
	return "RawStringLiteral"
}

func (s RawStringLiteral) GetContent() interface{} {
	return s.content
}

func (s RawStringLiteral) AsString() string {
	return s.content
}

func (s RawStringLiteral) accept(visitor *CypherRenderer) {
	visitor.enter(s)
	visitor.leave(s)
}

func (s RawStringLiteral) enter(renderer *CypherRenderer) {
	renderer.append(s.AsString())
}

func (s RawStringLiteral) leave(renderer *CypherRenderer) {
}
