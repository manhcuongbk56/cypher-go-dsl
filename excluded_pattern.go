package cypher_go_dsl

type ExcludedPattern struct {
	patternElement PatternElement
	key            string
	notNil         bool
	err            error
}

func ExcludedPatternCreate(patternElement PatternElement) ExcludedPattern {
	if patternElement.getError() != nil {
		return ExcludedPatternError(patternElement.getError())
	}
	e := ExcludedPattern{
		patternElement: patternElement,
		notNil:         true,
	}
	e.key = getAddress(&e)
	return e
}

func ExcludedPatternError(err error) ExcludedPattern {
	return ExcludedPattern{
		err: err,
	}
}

func (e ExcludedPattern) getError() error {
	return e.err
}

func (e ExcludedPattern) accept(visitor *CypherRenderer) {
	visitor.enter(e)
	NOT.accept(visitor)
	e.patternElement.accept(visitor)
	visitor.leave(e)
}

func (e ExcludedPattern) enter(renderer *CypherRenderer) {
}

func (e ExcludedPattern) leave(renderer *CypherRenderer) {
}

func (e ExcludedPattern) getKey() string {
	return e.key
}

func (e ExcludedPattern) isNotNil() bool {
	return e.notNil
}

func (e ExcludedPattern) GetExpressionType() ExpressionType {
	panic("implement me")
}

func (e ExcludedPattern) getConditionType() string {
	panic("implement me")
}
