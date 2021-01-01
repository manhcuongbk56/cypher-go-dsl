package cypher

import (
	"errors"
	"golang.org/x/xerrors"
)

type CompoundCondition struct {
	ConditionContainer
	operator      Operator
	conditions    []Condition
	conditionType ExpressionType
	key           string
	notNil        bool
	err           error
}

func CompoundConditionCreate(left Condition, operator Operator, right Condition) CompoundCondition {
	if left != nil && left.GetError() != nil {
		return CompoundConditionError(left.GetError())
	}
	if operator.GetError() != nil {
		return CompoundConditionError(operator.GetError())
	}
	if right != nil && right.GetError() != nil {
		return CompoundConditionError(right.GetError())
	}
	if left == nil || !left.isNotNil() {
		return CompoundConditionError(errors.New("left hand side condition is required"))
	}
	if !operator.isNotNil() {
		return CompoundConditionError(errors.New("operator is required"))
	}
	isOperatorValid := false
	for _, validOperator := range VALID_OPERATORS {
		if validOperator == operator {
			isOperatorValid = true
		}
	}
	if !isOperatorValid {
		return CompoundConditionError(xerrors.Errorf("operator %s is not valid", operator.representation))
	}
	if right == nil || !right.isNotNil() {
		return CompoundConditionError(errors.New("right hand side condition is required"))
	}
	condition := CompoundCondition{operator: operator, notNil: true}
	condition.add(operator, left)
	condition.add(operator, right)
	condition.injectKey()
	condition.ConditionContainer = ConditionWrap(condition)
	return condition
}

func CompoundConditionCreate1(operator Operator) CompoundCondition {
	condition := CompoundCondition{
		operator:   operator,
		conditions: make([]Condition, 0),
		notNil:     true,
	}
	condition.injectKey()
	condition.ConditionContainer = ConditionWrap(condition)
	return condition
}

func CompoundConditionError(err error) CompoundCondition {
	return CompoundCondition{
		err: err,
	}
}

func (c CompoundCondition) GetError() error {
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
			compound, isCompound := condition.(CompoundCondition)
			if isCompound {
				actualOperator = compound.operator
			} else {
				actualOperator = c.operator
			}
			acceptVisitorWithOperatorForChildCondition(visitor, actualOperator, condition)
		}
		visitor.leave(c)
	}
}

func (c CompoundCondition) enter(renderer *CypherRenderer) {
	renderer.append("(")
}

func (c CompoundCondition) leave(renderer *CypherRenderer) {
	renderer.append(")")
}

func (c CompoundCondition) getKey() string {
	return c.key
}

func (c CompoundCondition) GetExpressionType() ExpressionType {
	return c.conditionType
}

func (c CompoundCondition) Or(condition Condition) ConditionContainer {
	return ConditionWrap(c.add(OR, condition))
}

func (c CompoundCondition) And(condition Condition) ConditionContainer {
	return ConditionWrap(c.add(AND, condition))
}

func (c CompoundCondition) Xor(condition Condition) ConditionContainer {
	return ConditionWrap(c.add(XOR, condition))
}

func (c *CompoundCondition) injectKey() {
	c.key = getAddress(c)
}

var EMPTY_CONDITION = CompoundCondition{
	conditions:    make([]Condition, 0),
	conditionType: EMPTY_CONDITION_EXPRESSION,
	notNil:        true,
}

var VALID_OPERATORS = []Operator{AND, OR, XOR}

func (c *CompoundCondition) add(chainingOperator Operator, condition Condition) CompoundCondition {
	if c.GetExpressionType() == EMPTY_CONDITION_EXPRESSION {
		newCompound := CompoundConditionCreate1(chainingOperator)
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
	return CompoundConditionCreate(*c, chainingOperator, condition)
}

func (c CompoundCondition) hasCondition() bool {
	return !(c.GetExpressionType() == EMPTY_CONDITION_EXPRESSION ||
		len(c.conditions) == 0)
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
