package cypher

type ExposesWhere interface {
	Where(condition Condition) OngoingReadingWithWhere
	WhereConditionContainer(condition ConditionContainer) OngoingReadingWithWhere
	WherePattern(pattern RelationshipPattern) OngoingReadingWithWhere
}
