package cypher_go_dsl

type CompoundCondition struct {
	operator      Operator
	conditions    []Condition
	conditionType ExpressionType
	key string
	notNil        bool
	err error
}


func CompoundConditionCreate(left Expression, operator Operator, right Expression) CompoundCondition {
	condition := CompoundCondition{operator: operator}
	condition.injectKey()
	return condition
}

func CompoundConditionCreate1(operator Operator) CompoundCondition {
	condition := CompoundCondition{
		operator:   operator,
		conditions: make([]Condition, 0),
	}
	condition.injectKey()
	return condition
}


func (c CompoundCondition) getError() error {
	return c.err
}

func (c CompoundCondition) getConditionType() string {
	return "CompoundCondition"
}

func (c CompoundCondition) isNotNil() bool {
	return c.notNil
}

func (c CompoundCondition) accept(visitor *CypherRenderer) {
	if len(c.conditions) == 0 {
		return
	}
	hasManyConditions := len(c.conditions) > 1
	if hasManyConditions {
		visitor.enter(c)
	}
	acceptVisitorWithOperatorForChildCondition(visitor, Operator{}, c.conditions[0])

	if hasManyConditions {
		for _, condition := range c.conditions[1:] {
			var actualOperator Operator
			compound, isCompount := condition.(CompoundCondition)
			if isCompount {
				actualOperator = compound.operator
			}else {
				actualOperator = c.operator
			}
			acceptVisitorWithOperatorForChildCondition(visitor, actualOperator, condition)
		}
		visitor.leave(c)
	}
}

func (c CompoundCondition) enter(renderer *CypherRenderer) {
	panic("implement me")
}

func (c CompoundCondition) leave(renderer *CypherRenderer) {
	panic("implement me")
}

func (c CompoundCondition) getKey() string {
	return c.key
}

func (c CompoundCondition) GetExpressionType() ExpressionType {
	return c.conditionType
}

func (c *CompoundCondition) injectKey() {
	c.key = getAddress(c)
}

var EMPTY_CONDITION = CompoundCondition{
	conditions:    make([]Condition, 0),
	conditionType: EMPTY_CONDITION_EXPRESSION,
}

var VALID_OPERATORS = []Operator{AND, OR, XOR}



func (c *CompoundCondition) add(chainingOperator Operator, condition Condition) CompoundCondition {
	if c.GetExpressionType() == EMPTY_CONDITION_EXPRESSION {
		newCompound := CompoundCondition{
			operator: chainingOperator,
		}
		return newCompound.add(chainingOperator, condition)
	}
	if compoundCondition, isCompound := condition.(CompoundCondition); isCompound {
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
			inner := CompoundConditionCreate1(chainingOperator)
			inner.conditions = append(inner.conditions, compoundCondition)
			c.conditions = append(c.conditions, inner)
		}
		return *c
	}

	if c.operator == chainingOperator {
		c.conditions = append(c.conditions, condition)
		return *c
	}
	return CompoundConditionCreate(c, chainingOperator, condition)
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

func acceptVisitorWithOperatorForChildCondition(visitor *CypherRenderer, operator Operator, condition Condition) {
	VisitIfNotNull(operator, visitor)
	condition.accept(visitor)
}
