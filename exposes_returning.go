package cypher_go_dsl

type ExposesReturning interface {
	returningByString(variables ...string) OngoingReadingAndReturn
	returningByNamed(variables ...Named) OngoingReadingAndReturn
	returning(expression ...Expression) OngoingReadingAndReturn
	returningDistinctByString(variables ...string) OngoingReadingAndReturn
	returningDistinctByNamed(variables ...Named) OngoingReadingAndReturn
	returningDistinct(expression ...Expression) OngoingReadingAndReturn
}
