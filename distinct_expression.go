package cypher_go_dsl

type DistinctExpression struct {
	delegate Expression
	key      string
	notNil   bool
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
