package cypher_go_dsl

type ConditionBuilder struct {
	condition Expression
}

func (c *ConditionBuilder) Where(newCondition Expression)  {
	c.condition =  newCondition
}

func (c *ConditionBuilder) And(additionalCondition Expression)  {
	conditionContainer := ConditionContainer{ExpressionContainer{expression: additionalCondition}}
	conditionContainer.And(additionalCondition)
	c.condition = conditionContainer.expression
}

func (c *ConditionBuilder) Or(additionalCondition Expression)  {
	conditionContainer := ConditionContainer{ExpressionContainer{expression: additionalCondition}}
	conditionContainer.Or(additionalCondition)
	c.condition = conditionContainer.expression
}

func (c *ConditionBuilder) hasCondition() bool {
	return true
}

func (c *ConditionBuilder) buildCondition() Expression {
	return c.condition
}
