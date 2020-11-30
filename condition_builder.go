package cypher_go_dsl

type ConditionContainer struct {
	ExpressionContainer
}

func (b ConditionContainer) Where(condition *Expression) {
	b.condition = condition
}

func (b ConditionContainer) And(condition *Expression) {
	newCondition := (*(b.condition)).And(*condition)
	b.condition = &newCondition
}

func (b ConditionContainer) Or(condition *Expression) {
	newCondition := (*(b.condition)).Or(*condition)
	b.condition = &newCondition
}

func (b ConditionContainer) BuildCondition() *Expression {
	return b.condition
}
