package cypher

type NilLiteral struct {
	ExpressionContainer
	key    string
	notNil bool
	err    error
}

func NullLiteralCreate() NilLiteral {
	nilLiteral := NilLiteral{
		notNil: true,
	}
	nilLiteral.key = getAddress(&nilLiteral)
	nilLiteral.ExpressionContainer = ExpressionWrap(nilLiteral)
	return nilLiteral
}

var NIL_INSTANCE = NullLiteralCreate()

func (n NilLiteral) GetError() error {
	return n.err
}

func (n NilLiteral) accept(visitor *CypherRenderer) {
	visitor.enter(n)
	visitor.leave(n)
}

func (n NilLiteral) enter(renderer *CypherRenderer) {
	renderer.append(n.AsString())
}

func (n NilLiteral) leave(renderer *CypherRenderer) {
}

func (n NilLiteral) getKey() string {
	return n.key
}

func (n NilLiteral) isNotNil() bool {
	return n.notNil
}

func (n NilLiteral) GetExpressionType() ExpressionType {
	return "NilLiteral"
}

func (n NilLiteral) GetContent() interface{} {
	return nil
}

func (n NilLiteral) AsString() string {
	return "NULL"
}
