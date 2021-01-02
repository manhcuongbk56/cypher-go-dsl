package cypher

type ListPredicate struct {
	ExpressionContainer
	variable       SymbolicName
	listExpression Expression
	where          Where
	key            string
	notNil         bool
	err            error
}

func ListPredicateCreate(variable SymbolicName, listExpression Expression, where Where) ListPredicate {
	if variable.GetError() != nil {
		return ListPredicateError(variable.GetError())
	}
	if listExpression != nil && listExpression.GetError() != nil {
		return ListPredicateError(variable.GetError())
	}
	if where.GetError() != nil {
		return ListPredicateError(variable.GetError())
	}
	listPredicate := ListPredicate{
		variable:       variable,
		listExpression: listExpression,
		where:          where,
		notNil:         true,
	}
	listPredicate.key = getAddress(&listPredicate)
	listPredicate.ExpressionContainer = ExpressionWrap(listPredicate)
	return listPredicate
}

func ListPredicateError(err error) ListPredicate {
	return ListPredicate{err: err}
}

func (l ListPredicate) GetError() error {
	return l.err
}

func (l ListPredicate) accept(visitor *CypherRenderer) {
	visitor.enter(l)
	l.variable.accept(visitor)
	IN.accept(visitor)
	l.listExpression.accept(visitor)
	l.where.accept(visitor)
	visitor.leave(l)
}

func (l ListPredicate) enter(renderer *CypherRenderer) {
}

func (l ListPredicate) leave(renderer *CypherRenderer) {
}

func (l ListPredicate) getKey() string {
	return l.key
}

func (l ListPredicate) isNotNil() bool {
	return l.notNil
}

func (l ListPredicate) GetExpressionType() ExpressionType {
	return "ListPredicate"
}
