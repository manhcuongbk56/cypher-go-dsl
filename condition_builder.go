package cypher_go_dsl

type ConditionBuilder struct {
	condition *Expression
}

func (b ConditionBuilder) Where(condition *Expression) {
	b.condition = condition
}

func (b ConditionBuilder) And(condition *Expression) {
	newCondition := (*(b.condition)).And(*condition)
	b.condition = &newCondition
}

func (b ConditionBuilder) Or(condition *Expression) {
	newCondition := (*(b.condition)).Or(*condition)
	b.condition = &newCondition
}

func (b ConditionBuilder) BuildCondition() *Condition {
	return b.condition
}
