package cypher

type RelationshipPatternCondition struct {
	ConditionContainer
	pathPattern RelationshipPattern
	key         string
	notNil      bool
	err         error
}

func RelationshipPatternConditionCreate(pathPattern RelationshipPattern) RelationshipPatternCondition {
	if pathPattern.getError() != nil {
		return RelationshipPatternConditionError(pathPattern.getError())
	}
	r := RelationshipPatternCondition{
		pathPattern: pathPattern,
		notNil:      true,
	}
	r.key = getAddress(&r)
	r.ConditionContainer = ConditionWrap(r)
	return r
}

func RelationshipPatternConditionError(err error) RelationshipPatternCondition {
	return RelationshipPatternCondition{
		err: err,
	}
}

func (r RelationshipPatternCondition) getConditionType() string {
	return "RelationshipPatternCondition"
}

func (r RelationshipPatternCondition) getError() error {
	return r.err
}

func (r RelationshipPatternCondition) isNotNil() bool {
	return r.notNil
}

func (r RelationshipPatternCondition) accept(visitor *CypherRenderer) {
	visitor.enter(r)
	r.pathPattern.accept(visitor)
	visitor.leave(r)
}

func (r RelationshipPatternCondition) enter(renderer *CypherRenderer) {
}

func (r RelationshipPatternCondition) leave(renderer *CypherRenderer) {
}

func (r RelationshipPatternCondition) getKey() string {
	return r.key
}

func (r RelationshipPatternCondition) GetExpressionType() ExpressionType {
	return "RelationshipPatternCondition"
}
