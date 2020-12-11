package cypher_go_dsl

type RelationshipPatternCondition struct {
	pathPattern RelationshipPattern
	key         string
	notNil      bool
	err         error
}

func RelationshipPatternConditionCreate(pathPattern RelationshipPattern) RelationshipPatternCondition {
	r := RelationshipPatternCondition{
		pathPattern: pathPattern,
		notNil:      true,
	}
	r.key = getAddress(&r)
	return r
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
	panic("implement me")
}

func (r RelationshipPatternCondition) GetExpressionType() ExpressionType {
	panic("implement me")
}
