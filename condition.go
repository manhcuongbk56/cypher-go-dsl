package cypher

type Condition interface {
	Expression
	getConditionType() string
	And(condition Condition) ConditionContainer
	Or(condition Condition) ConditionContainer
	Xor(condition Condition) ConditionContainer
	AndPattern(pathPattern RelationshipPattern) ConditionContainer
	OrPattern(pathPattern RelationshipPattern) ConditionContainer
	XorPattern(pathPattern RelationshipPattern) ConditionContainer
	Not() ConditionContainer
}
