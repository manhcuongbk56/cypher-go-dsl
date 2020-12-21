package cypher

type Arguments struct {
	expressions []Expression
	key         string
	notNil      bool
	err         error
}

func ArgumentsCreate(expression []Expression) Arguments {
	for _, expr := range expression {
		if expr != nil && expr.getError() != nil {
			return Arguments{err: expr.getError()}
		}
	}
	a := Arguments{
		expressions: expression,
		notNil:      true,
	}
	a.key = getAddress(&a)
	return a
}

func (a Arguments) getError() error {
	return a.err
}

func (a Arguments) isNotNil() bool {
	return a.notNil
}

func (a Arguments) getKey() string {
	return a.key
}

func (a Arguments) PrepareVisit(child Visitable) Visitable {
	expression, isExpression := child.(Expression)
	if !isExpression {
		panic("Can not prepare un expression type in expression list")
	}
	return NameOrExpression(expression)
}

func (a Arguments) accept(visitor *CypherRenderer) {
	(*visitor).enter(a)
	for _, expression := range a.expressions {
		a.PrepareVisit(expression).accept(visitor)
	}
	(*visitor).leave(a)
}

func (a Arguments) enter(renderer *CypherRenderer) {
}

func (a Arguments) leave(renderer *CypherRenderer) {
}
