package cypher_go_dsl

type ConditionBuilder struct {
	condition *Condition
}

func (b ConditionBuilder) Where(condition *Condition) {
	b.condition = condition
}

func (b ConditionBuilder) And(condition *Condition) {
	newCondition := (*(b.condition)).And(*condition)
	b.condition = &newCondition
}

func (b ConditionBuilder) Or(condition *Condition) {
	newCondition := (*(b.condition)).Or(*condition)
	b.condition = &newCondition
}

func (b ConditionBuilder) BuildCondition() *Condition {
	return b.condition
}
