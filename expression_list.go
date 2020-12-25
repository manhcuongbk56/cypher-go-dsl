package cypher

type ExpressionList struct {
	ExpressionContainer
	expressions []Expression
	key         string
	notNil      bool
	err         error
}

func ExpressionListCreate(expressions []Expression) ExpressionList {
	for _, expression := range expressions {
		if expression.getError() != nil {
			return ExpressionListError(expression.getError())
		}
	}
	e := ExpressionList{
		expressions: expressions,
		notNil:      true,
	}
	e.key = getAddress(&e)
	e.ExpressionContainer = ExpressionWrap(e)
	return e
}

func ExpressionListError(err error) ExpressionList {
	return ExpressionList{
		err: err,
	}
}

func (e ExpressionList) GetExpressionType() ExpressionType {
	return "ExpressionList"
}

func (e ExpressionList) getError() error {
	return e.err
}

func (e ExpressionList) isNotNil() bool {
	return e.notNil
}

func (e ExpressionList) getKey() string {
	return e.key
}

func (e ExpressionList) PrepareVisit(child Visitable) Visitable {
	expression, isExpression := child.(Expression)
	if !isExpression {
		panic("Can not prepare un expression type in expression list")
	}
	return NameOrExpression(expression)
}

func (e ExpressionList) accept(visitor *CypherRenderer) {
	visitor.enter(e)
	for _, expression := range e.expressions {
		e.PrepareVisit(expression).accept(visitor)
	}
	visitor.leave(e)
}

func (e ExpressionList) enter(renderer *CypherRenderer) {
}

func (e ExpressionList) leave(renderer *CypherRenderer) {
}

func NewExpressionList(expression ...Expression) ExpressionList {
	expressions := make([]Expression, len(expression))
	for i := range expression {
		expressions[i] = expression[i]
	}
	return ExpressionList{expressions: expressions}
}
