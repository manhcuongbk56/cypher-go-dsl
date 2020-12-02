package cypher_go_dsl

type Condition interface {
	Expression
	getConditionType() string
}
