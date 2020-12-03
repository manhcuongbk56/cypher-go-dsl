package cypher_go_dsl

type ConditionContainer struct {
	ExpressionContainer
}

func (b *ConditionContainer) And(condition Condition) *ConditionContainer {
	b.expression = CompoundConditionCreate(b.expression, AND, condition)
	return b
}

func (b *ConditionContainer) Or(condition Condition) *ConditionContainer {
	b.expression = CompoundConditionCreate(b.expression, OR, condition)
	return b
}

func (b *ConditionContainer) Xor(condition Condition) *ConditionContainer {
	b.expression = CompoundConditionCreate(b.expression, XOR, condition)
	return b
}

func (b *ConditionContainer) AndRelationshipPattern(pathPattern RelationshipPattern) *ConditionContainer {
	b.expression = CompoundConditionCreate(b.expression, AND, RelationshipPatternCondition{
		pathPattern: pathPattern,
	})
	return b
}

func (b *ConditionContainer) OrRelationshipPattern(pathPattern RelationshipPattern) *ConditionContainer {
	b.expression = CompoundConditionCreate(b.expression, OR, RelationshipPatternCondition{
		pathPattern: pathPattern,
	})
	return b
}

func (b *ConditionContainer) XorRelationshipPattern(pathPattern RelationshipPattern) *ConditionContainer {
	b.expression = CompoundConditionCreate(b.expression, XOR, RelationshipPatternCondition{
		pathPattern: pathPattern,
	})
	return b
}

func (b *ConditionContainer) Not(pathPattern RelationshipPattern) *ConditionContainer {
	b.expression = NewComparisonWithConstant(NOT, b.expression)
	return b
}
