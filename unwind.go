package cypher

type unwind struct {
	expressionToUnwind Expression
	variable           string
	key                string
	notNil             bool
	err                error
}

func unwindCreate(expressionToUnwind Expression, variable string) unwind {
	if expressionToUnwind != nil && expressionToUnwind.GetError() != nil {
		return unwindError(expressionToUnwind.GetError())
	}
	var expression Expression
	if aliased, isAliased := expressionToUnwind.(Aliased); isAliased {
		expression = aliased.AsName()
	} else {
		expression = expressionToUnwind
	}
	unwind := unwind{
		expressionToUnwind: expression,
		variable:           variable,
		notNil:             true,
	}
	unwind.key = getAddress(&unwind)
	return unwind
}

func unwindError(err error) unwind {
	return unwind{err: err}
}

func (u unwind) GetError() error {
	return u.err
}

func (u unwind) accept(visitor *CypherRenderer) {
	visitor.enter(u)
	u.expressionToUnwind.accept(visitor)
	visitor.leave(u)
}

func (u unwind) enter(renderer *CypherRenderer) {
	renderer.append("UNWIND ")
}

func (u unwind) leave(renderer *CypherRenderer) {
	renderer.append(" AS ")
	renderer.append(u.variable)
	renderer.append(" ")
}

func (u unwind) getKey() string {
	return u.key
}

func (u unwind) isNotNil() bool {
	return u.notNil
}
