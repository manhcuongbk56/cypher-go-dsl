package cypher_go_dsl

type ExposesWhere interface {
	Where(condition Condition) OngoingReadingWithWhere
	WherePattern(pattern RelationshipPattern) OngoingReadingWithWhere
}
