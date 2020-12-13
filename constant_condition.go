package cypher_go_dsl

type ConstantCondition struct {
	value  BooleanLiteral
	key    string
	notNil bool
	err    error
}

var TRUE_CONDITION = ConstantConditionCreate(TRUE)
var FALSE_CONDITION = ConstantConditionCreate(FALSE)

func ConstantConditionCreate(value BooleanLiteral) ConstantCondition {
	if value.getError() != nil {
		return ConstantConditionError(value.getError())
	}
	constantCondition := ConstantCondition{
		value:  value,
		notNil: true,
	}
	constantCondition.key = getAddress(&constantCondition)
	return constantCondition
}

func ConstantConditionError(err error) ConstantCondition {
	return ConstantCondition{err: err}
}

func (c ConstantCondition) getError() error {
	return c.err
}

func (c ConstantCondition) accept(visitor *CypherRenderer) {
	visitor.enter(c)
	c.value.accept(visitor)
	visitor.leave(c)
}

func (c ConstantCondition) enter(renderer *CypherRenderer) {

}

func (c ConstantCondition) leave(renderer *CypherRenderer) {

}

func (c ConstantCondition) getKey() string {
	return c.key
}

func (c ConstantCondition) isNotNil() bool {
	return c.notNil
}

func (c ConstantCondition) GetExpressionType() ExpressionType {
	return "ConstantCondition"
}

func (c ConstantCondition) getConditionType() string {
	return "ConstantCondition"
}
