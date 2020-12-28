package cypher

type NestedExpression struct {
	ExpressionContainer
	delegate Expression
	key      string
	notNil   bool
	err      error
}

func NestedExpressionCreate(delegate Expression) NestedExpression {
	if delegate != nil && delegate.getError() != nil {
		return NestedExpressionError(delegate.getError())
	}
	nested := NestedExpression{
		delegate: delegate,
		notNil:   true,
	}
	nested.key = getAddress(&nested)
	nested.ExpressionContainer = ExpressionWrap(nested)
	return nested
}

func NestedExpressionError(err error) NestedExpression {
	return NestedExpression{
		err: err,
	}
}

func (n NestedExpression) getError() error {
	return n.err
}

func (n NestedExpression) accept(visitor *CypherRenderer) {
	visitor.enter(n)
	NameOrExpression(n.delegate).accept(visitor)
	visitor.leave(n)
}

func (n NestedExpression) enter(renderer *CypherRenderer) {
	renderer.append("(")
}

func (n NestedExpression) leave(renderer *CypherRenderer) {
	renderer.append(")")
}

func (n NestedExpression) getKey() string {
	return n.key
}

func (n NestedExpression) isNotNil() bool {
	return n.notNil
}

func (n NestedExpression) GetExpressionType() ExpressionType {
	return "NestedExpression"
}
