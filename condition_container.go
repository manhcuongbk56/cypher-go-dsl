package cypher

type ConditionContainer struct {
	ExpressionContainer
}

func (b *ConditionContainer) And(condition Condition) *ConditionContainer {
	b.expression = CompoundConditionCreate(b.expression.(Condition), AND, condition)
	return b
}

func (b *ConditionContainer) Or(condition Condition) *ConditionContainer {
	b.expression = CompoundConditionCreate(b.expression.(Condition), OR, condition)
	return b
}

func (b *ConditionContainer) Xor(condition Condition) *ConditionContainer {
	b.expression = CompoundConditionCreate(b.expression.(Condition), XOR, condition)
	return b
}

func (b *ConditionContainer) AndRelationshipPattern(pathPattern RelationshipPattern) *ConditionContainer {
	b.expression = CompoundConditionCreate(b.expression.(Condition), AND, RelationshipPatternCondition{
		pathPattern: pathPattern,
	})
	return b
}

func (b *ConditionContainer) OrRelationshipPattern(pathPattern RelationshipPattern) *ConditionContainer {
	b.expression = CompoundConditionCreate(b.expression.(Condition), OR, RelationshipPatternCondition{
		pathPattern: pathPattern,
	})
	return b
}

func (b *ConditionContainer) XorRelationshipPattern(pathPattern RelationshipPattern) *ConditionContainer {
	b.expression = CompoundConditionCreate(b.expression.(Condition), XOR, RelationshipPatternCondition{
		pathPattern: pathPattern,
	})
	return b
}

func (b *ConditionContainer) Not(pathPattern RelationshipPattern) *ConditionContainer {
	b.expression = ComparisonCreate1(NOT, b.expression)
	return b
}
