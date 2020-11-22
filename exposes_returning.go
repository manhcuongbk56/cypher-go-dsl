package cypher_go_dsl

type ExposesReturning interface {
	returning(expression ...IsExpression) OngoingReadingAndReturn
	returningDistinct(expression ...IsExpression) OngoingReadingAndReturn
}