package cypher_go_dsl

type CompoundCondition struct {
	operator Operator
	conditions []Expression
	conditionType ExpressionType
}

func (c CompoundCondition) accept(visitor *CypherRenderer) {
	panic("implement me")
}

func (c CompoundCondition) enter(renderer *CypherRenderer) {
	panic("implement me")
}

func (c CompoundCondition) leave(renderer *CypherRenderer) {
	panic("implement me")
}

func (c CompoundCondition) getKey() string {
	panic("implement me")
}

func (c CompoundCondition) GetExpressionType() ExpressionType {
	return c.conditionType
}

var EMPTY_CONDITION = CompoundCondition{
	conditions: make([]Expression, 0),
	conditionType: EMPTY_CONDITION_EXPRESSION,
}

var VALID_OPERATORS = []Operator{AND, OR, XOR}

func (c CompoundCondition) CompoundConditionCreate(left Expression, operator Operator, right Expression) {
	return CompoundCondition{operator: operator}
}

func (c *CompoundCondition) add(chainingOperator Operator, expression Expression) CompoundCondition {
	if c.GetExpressionType() == EMPTY_CONDITION_EXPRESSION {
		newCompound := CompoundCondition{
			operator: chainingOperator,
		}
		return newCompound.add(chainingOperator, expression)
	}
	if compoundCondition, isCompound := expression.(CompoundCondition); isCompound  {
		if !compoundCondition.hasCondition() {
			return c
		}
		if c.operator == chainingOperator && chainingOperator == compoundCondition.operator {
			c.conditions = append(c.conditions, compoundCondition.conditions...)
		}
	}

}

func (c CompoundCondition) hasCondition() bool {
	return !(c.GetExpressionType() == EMPTY_CONDITION_EXPRESSION ||
		len(c.conditions) > 0)
}
