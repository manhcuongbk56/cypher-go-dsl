package cypher_go_dsl

import "fmt"

type RelationshipPatternCondition struct {
	pathPattern RelationshipPattern
	key         string
	notNil      bool
}

func (r RelationshipPatternCondition) isNotNil() bool {
	return r.notNil
}

func (r RelationshipPatternCondition) accept(visitor *CypherRenderer) {
	r.key = fmt.Sprint(&r)
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
