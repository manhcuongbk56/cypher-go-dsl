package cypher_go_dsl

type ExpressionList struct {
	expressions []IsExpression
}

func (e ExpressionList) PrepareVisit(child Visitable) Visitable {
	expression, isExpression := child.(IsExpression)
	if !isExpression {
		panic("Can not prepare un expression type in expression list")
	}
	return NameOrExpression(expression)
}

func (e ExpressionList) Accept(visitor Visitor) {
	visitor.Enter(e)
	for _, expression := range e.expressions{
		e.PrepareVisit(expression).Accept(visitor)
	}
	visitor.Leave(e)
}

func NewExpressionList(expression ...IsExpression) ExpressionList {
	expressions := make([]IsExpression, len(expression))
	expressions = append(expressions, expression...)
	return ExpressionList{expressions: expressions}
}

