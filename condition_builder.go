package cypher

type ConditionBuilder struct {
	condition Condition
}

func ConditionBuilderCreate() ConditionBuilder {
	return ConditionBuilder{}
}

func (c *ConditionBuilder) Where(newCondition Condition) {
	c.condition = newCondition
}

func (c *ConditionBuilder) And(additionalCondition Condition) {
	conditionContainer := ConditionContainer{ExpressionContainer{expression: additionalCondition}}
	conditionContainer.And(additionalCondition)
	c.condition = conditionContainer.expression.(Condition)
}

func (c *ConditionBuilder) Or(additionalCondition Condition) {
	conditionContainer := ConditionContainer{ExpressionContainer{expression: additionalCondition}}
	conditionContainer.Or(additionalCondition)
	c.condition = conditionContainer.expression.(Condition)
}

func (c *ConditionBuilder) hasCondition() bool {
	return true
}

func (c *ConditionBuilder) buildCondition() Condition {
	return c.condition
}
