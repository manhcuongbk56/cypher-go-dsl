package cypher_go_dsl

type ExposesReturning interface {
	returning(expression ...Expression) OngoingReadingAndReturn
	returningDistinct(expression ...Expression) OngoingReadingAndReturn
}