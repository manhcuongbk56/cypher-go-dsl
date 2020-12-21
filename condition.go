package cypher

type Condition interface {
	Expression
	getConditionType() string
}
