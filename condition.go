package cypher

type Condition interface {
	Expression
	getConditionType() string
	And(condition Condition) ConditionContainer
	Or(condition Condition) ConditionContainer
	Xor(condition Condition) ConditionContainer
	AndRelationshipPattern(pathPattern RelationshipPattern) ConditionContainer
	OrRelationshipPattern(pathPattern RelationshipPattern) ConditionContainer
	XorRelationshipPattern(pathPattern RelationshipPattern) ConditionContainer
	Not() ConditionContainer
}
