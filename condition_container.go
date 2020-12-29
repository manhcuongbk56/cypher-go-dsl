package cypher

type ConditionContainer struct {
	ExpressionContainer
}

func ConditionWrap(condition Condition) ConditionContainer {
	return ConditionContainer{
		ExpressionContainer{condition},
	}
}

func (b ConditionContainer) And(condition Condition) ConditionContainer {
	if compoundCondition, isCompound := b.expression.(CompoundCondition); isCompound {
		return compoundCondition.And(condition)
	}
	b.expression = CompoundConditionCreate(b.expression.(Condition), AND, condition)
	return b
}

func (b ConditionContainer) Or(condition Condition) ConditionContainer {
	if compoundCondition, isCompound := b.expression.(CompoundCondition); isCompound {
		return compoundCondition.Or(condition)
	}
	b.expression = CompoundConditionCreate(b.expression.(Condition), OR, condition)
	return b
}

func (b ConditionContainer) Xor(condition Condition) ConditionContainer {
	if compoundCondition, isCompound := b.expression.(CompoundCondition); isCompound {
		return compoundCondition.Xor(condition)
	}
	b.expression = CompoundConditionCreate(b.expression.(Condition), XOR, condition)
	return b
}

func (b ConditionContainer) AndPattern(pathPattern RelationshipPattern) ConditionContainer {
	b.expression = CompoundConditionCreate(b.expression.(Condition), AND, RelationshipPatternConditionCreate(pathPattern))
	return b
}

func (b ConditionContainer) OrPattern(pathPattern RelationshipPattern) ConditionContainer {
	b.expression = CompoundConditionCreate(b.expression.(Condition), OR, RelationshipPatternConditionCreate(pathPattern))
	return b
}

func (b ConditionContainer) XorPattern(pathPattern RelationshipPattern) ConditionContainer {
	b.expression = CompoundConditionCreate(b.expression.(Condition), XOR, RelationshipPatternConditionCreate(pathPattern))
	return b
}

func (b ConditionContainer) Not() ConditionContainer {
	b.expression = ComparisonCreate1(NOT, b.expression)
	return b
}

func (b ConditionContainer) Get() Condition {
	return b.expression.(Condition)
}
