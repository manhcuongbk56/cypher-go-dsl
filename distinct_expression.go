package cypher

type DistinctExpression struct {
	ExpressionContainer
	delegate Expression
	key      string
	notNil   bool
	err      error
}

func DistinctExpressionCreate(delegate Expression) DistinctExpression {
	if delegate != nil && delegate.GetError() != nil {
		return DistinctExpressionError(delegate.GetError())
	}
	d := DistinctExpression{
		delegate: delegate,
		notNil:   true,
	}
	d.key = getAddress(&d)
	d.ExpressionContainer = ExpressionWrap(d)
	return d
}

func DistinctExpressionError(err error) DistinctExpression {
	return DistinctExpression{err: err}
}

func (d DistinctExpression) GetError() error {
	return d.err
}

func (d DistinctExpression) isNotNil() bool {
	return d.notNil
}

func (d DistinctExpression) accept(visitor *CypherRenderer) {
	visitor.enter(d)
	distinct := Distinct{
		IsDistinct: false,
	}
	distinct.accept(visitor)
	d.delegate.accept(visitor)
	visitor.leave(d)
}

func (d DistinctExpression) enter(renderer *CypherRenderer) {
}

func (d DistinctExpression) leave(renderer *CypherRenderer) {
}

func (d DistinctExpression) getKey() string {
	return d.key
}

func (d DistinctExpression) GetExpressionType() ExpressionType {
	return EXPRESSION
}
