package cypher_go_dsl

type CompoundCondition struct {
	operator      Operator
	conditions    []Expression
	conditionType ExpressionType
	notNil        bool
}

func (c CompoundCondition) isNotNil() bool {
	return c.notNil
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
	conditions:    make([]Expression, 0),
	conditionType: EMPTY_CONDITION_EXPRESSION,
}

var VALID_OPERATORS = []Operator{AND, OR, XOR}

func CompoundConditionCreate(left Expression, operator Operator, right Expression) CompoundCondition {
	return CompoundCondition{operator: operator}
}

func CompoundConditionCreateWithOperator(operator Operator) CompoundCondition {
	return CompoundCondition{
		operator:   operator,
		conditions: make([]Expression, 0),
	}
}

func (c *CompoundCondition) add(chainingOperator Operator, expression Expression) CompoundCondition {
	if c.GetExpressionType() == EMPTY_CONDITION_EXPRESSION {
		newCompound := CompoundCondition{
			operator: chainingOperator,
		}
		return newCompound.add(chainingOperator, expression)
	}
	if compoundCondition, isCompound := expression.(CompoundCondition); isCompound {
		if !compoundCondition.hasCondition() {
			return *c
		}
		if c.operator == chainingOperator && chainingOperator == compoundCondition.operator {
			if c.canBeFlattenedWith(chainingOperator) {
				c.conditions = append(c.conditions, compoundCondition.conditions...)
			} else {
				c.conditions = append(c.conditions, compoundCondition)
			}
		} else {
			inner := CompoundConditionCreateWithOperator(chainingOperator)
			inner.conditions = append(inner.conditions, compoundCondition)
			c.conditions = append(c.conditions, inner)
		}
		return *c
	}

	if c.operator == chainingOperator {
		c.conditions = append(c.conditions, expression)
		return *c
	}
	return CompoundConditionCreate(c, chainingOperator, expression)
}

func (c CompoundCondition) hasCondition() bool {
	return !(c.GetExpressionType() == EMPTY_CONDITION_EXPRESSION ||
		len(c.conditions) > 0)
}

func (c CompoundCondition) canBeFlattenedWith(operator Operator) bool {
	for _, c := range c.conditions {
		if compound, isCompound := c.(CompoundCondition); isCompound && compound.operator == operator {
			return false
		}
	}
	return true
}
