package cypher_go_dsl

type ExposesReturning interface {
	ReturningByString(variables ...string) OngoingReadingAndReturn
	ReturningByNamed(variables ...Named) OngoingReadingAndReturn
	Returning(expression ...Expression) OngoingReadingAndReturn
	ReturningDistinctByString(variables ...string) OngoingReadingAndReturn
	ReturningDistinctByNamed(variables ...Named) OngoingReadingAndReturn
	ReturningDistinct(expression ...Expression) OngoingReadingAndReturn
}
