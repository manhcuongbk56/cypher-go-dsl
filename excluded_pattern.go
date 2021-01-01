package cypher

type ExcludedPattern struct {
	ConditionContainer
	patternElement PatternElement
	key            string
	notNil         bool
	err            error
}

func ExcludedPatternCreate(patternElement PatternElement) ExcludedPattern {
	if patternElement.GetError() != nil {
		return ExcludedPatternError(patternElement.GetError())
	}
	e := ExcludedPattern{
		patternElement: patternElement,
		notNil:         true,
	}
	e.key = getAddress(&e)
	e.ConditionContainer = ConditionWrap(e)
	return e
}

func ExcludedPatternError(err error) ExcludedPattern {
	return ExcludedPattern{
		err: err,
	}
}

func (e ExcludedPattern) GetError() error {
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
