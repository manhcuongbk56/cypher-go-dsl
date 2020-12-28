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
	c.condition = c.condition.And(additionalCondition).Get()
}

func (c *ConditionBuilder) Or(additionalCondition Condition) {
	c.condition = c.condition.Or(additionalCondition).Get()
}

func (c *ConditionBuilder) hasCondition() bool {
	if c.condition == nil || !c.condition.isNotNil() {
		return false
	}
	compound, isCompound := c.condition.(CompoundCondition)
	return c.condition.isNotNil() && (!isCompound || compound.hasCondition())
}

func (c *ConditionBuilder) buildCondition() Condition {
	if c.hasCondition() {
		return c.condition
	}
	return nil
}
