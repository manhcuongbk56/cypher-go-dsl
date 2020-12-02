package cypher_go_dsl

type ExposesWhere interface {
	where(condition Condition) OngoingReadingWithWhere
	wherePattern(pattern RelationshipPattern) OngoingReadingWithWhere
}
