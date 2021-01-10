package cypher

type UnwindPhrase struct {
	expressionToUnwind Expression
	variable           string
	key                string
	notNil             bool
	err                error
}

func unwindCreate(expressionToUnwind Expression, variable string) UnwindPhrase {
	if expressionToUnwind != nil && expressionToUnwind.GetError() != nil {
		return unwindError(expressionToUnwind.GetError())
	}
	var expression Expression
	if aliased, isAliased := expressionToUnwind.(Aliased); isAliased {
		expression = aliased.AsName()
	} else {
		expression = expressionToUnwind
	}
	unwind := UnwindPhrase{
		expressionToUnwind: expression,
		variable:           variable,
		notNil:             true,
	}
	unwind.key = getAddress(&unwind)
	return unwind
}

func unwindError(err error) UnwindPhrase {
	return UnwindPhrase{err: err}
}

func (u UnwindPhrase) GetError() error {
	return u.err
}

func (u UnwindPhrase) accept(visitor *CypherRenderer) {
	visitor.enter(u)
	u.expressionToUnwind.accept(visitor)
	visitor.leave(u)
}

func (u UnwindPhrase) enter(renderer *CypherRenderer) {
	renderer.append("UNWIND ")
}

func (u UnwindPhrase) leave(renderer *CypherRenderer) {
	renderer.append(" AS ")
	renderer.append(u.variable)
	renderer.append(" ")
}

func (u UnwindPhrase) getKey() string {
	return u.key
}

func (u UnwindPhrase) isNotNil() bool {
	return u.notNil
}
